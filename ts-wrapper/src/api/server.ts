import { config as loadEnvFile } from 'dotenv';
import { MongoClient } from 'mongodb';
import { createServer } from 'node:http';
import { resolve } from 'path';

loadEnvFile({ path: resolve(process.cwd(), '../enviromments/production/stops/.env') });

const port = Number(process.env.PORT || 3000);
const databaseUri = process.env.DATABASE_URI || '';
const databaseName = process.env.DATABASE_NAME || '';
const stopsCollection = process.env.STOPS_COLLECTION || 'stops';
const cacheTtlMs = Number(process.env.STOPS_CACHE_TTL_MS || 30000);

let mongoClient: MongoClient | null = null;
let cachedStopIds: string[] = [];
let cachedAt = 0;
let lastFetchAt = 0;
let lastFetchError: null | string = null;

function emitDebugLog(hypothesisId: string, location: string, message: string, data: Record<string, unknown>) {
	const payload = {
		data,
		hypothesisId,
		location,
		message,
		runId: 'run-3',
		timestamp: Date.now(),
	};
	fetch('http://127.0.0.1:7242/ingest/637f32fa-3093-410b-9cee-497ee5d799e7', { body: JSON.stringify(payload), headers: { 'Content-Type': 'application/json' }, method: 'POST' }).catch(() => {});
}

async function getMongoClient(): Promise<MongoClient> {
	if (mongoClient) {
		return mongoClient;
	}

	if (!databaseUri) {
		throw new Error('DATABASE_URI is required');
	}

	mongoClient = new MongoClient(databaseUri);
	await mongoClient.connect();
	return mongoClient;
}

async function fetchStopIds(): Promise<string[]> {
	const now = Date.now();
	// #region agent log
	emitDebugLog('H1', 'ts-wrapper/src/api/server.ts:fetchStopIds:entry', 'Evaluating stop_id fetch preconditions', {
		cachedSize: cachedStopIds.length,
		cacheTtlMs,
		databaseNameConfigured: databaseName.length > 0,
		databaseUriConfigured: databaseUri.length > 0,
		stopsCollection,
	});
	// #endregion
	if (now - cachedAt <= cacheTtlMs && cachedStopIds.length > 0) {
		lastFetchError = null;
		lastFetchAt = now;
		return cachedStopIds;
	}

	if (!databaseName) {
		throw new Error('DATABASE_NAME is required');
	}

	const client = await getMongoClient();
	const collection = client.db(databaseName).collection(stopsCollection);
	let ids: string[] = [];
	try {
		const docs = await collection.aggregate<{ stop_id: string }>([
			{
				$project: {
					_id: 0,
					stop_id: { $trim: { input: '$stop_id' } },
				},
			},
			{
				$match: {
					stop_id: { $ne: '' },
				},
			},
			{
				$group: {
					_id: '$stop_id',
				},
			},
			{
				$project: {
					_id: 0,
					stop_id: '$_id',
				},
			},
			{
				$sort: {
					stop_id: 1,
				},
			},
		]).toArray();

		ids = docs
			.map(item => item.stop_id)
			.filter((value): value is string => typeof value === 'string' && value.length > 0);
		lastFetchError = null;
	} catch (error) {
		const message = error instanceof Error ? error.message : 'Unknown error';
		const isUnauthorizedAggregate = message.toLowerCase().includes('not authorized') && message.toLowerCase().includes('aggregate');

		// #region agent log
		emitDebugLog('H6', 'ts-wrapper/src/api/server.ts:fetchStopIds:aggregate-fallback', 'Aggregate query failed; evaluating fallback', {
			fallbackToFind: isUnauthorizedAggregate,
		});
		// #endregion

		if (!isUnauthorizedAggregate) {
			lastFetchError = message;
			throw error;
		}

		try {
			const docs = await collection.find(
				{},
				{ projection: { _id: 0, stop_id: 1 } },
			).toArray();

			ids = docs
				.map(item => item?.stop_id)
				.filter((value): value is string => typeof value === 'string' && value.trim().length > 0)
				.map(value => value.trim());
			lastFetchError = null;
		} catch (fallbackError) {
			lastFetchError = fallbackError instanceof Error ? fallbackError.message : 'Unknown error';
			throw fallbackError;
		}
	}

	cachedStopIds = Array.from(new Set(ids)).sort();
	cachedAt = now;
	lastFetchAt = now;
	// #region agent log
	emitDebugLog('H2', 'ts-wrapper/src/api/server.ts:fetchStopIds:aggregation', 'Mongo aggregation completed for stop_ids', {
		aggregatedCount: cachedStopIds.length,
		stopsCollection,
	});
	// #endregion
	return cachedStopIds;
}

const server = createServer(async (req, res) => {
	if (!req.url) {
		res.statusCode = 400;
		res.end('Bad Request');
		return;
	}
	const requestUrl = new URL(req.url, 'http://localhost');
	const path = requestUrl.pathname;

	// #region agent log
	emitDebugLog('H4', 'ts-wrapper/src/api/server.ts:request', 'Incoming API request', {
		method: req.method || '',
		path,
		url: req.url,
	});
	// #endregion

	if (req.method === 'GET' && (path === '/health' || path === '/health/')) {
		res.setHeader('Content-Type', 'application/json');
		res.end(JSON.stringify({
			cacheSize: cachedStopIds.length,
			databaseConfigured: databaseUri.length > 0 && databaseName.length > 0,
			lastFetchAt,
			lastFetchError,
			status: 'ok',
		}));
		return;
	}

	if (req.method === 'GET' && (path === '/reference/stops' || path === '/reference/stops/' || path === '/stops' || path === '/stops/')) {
		try {
			const stopIds = await fetchStopIds();
			res.setHeader('Content-Type', 'application/json');
			res.end(JSON.stringify({
				count: stopIds.length,
				stop_ids: stopIds,
			}));
		} catch (error) {
			const errorMessage = error instanceof Error ? error.message : 'Unknown error';
			const isUnauthorized = errorMessage.toLowerCase().includes('not authorized');
			// #region agent log
			emitDebugLog('H3', 'ts-wrapper/src/api/server.ts:/reference/stops:catch', 'Failed to serve /reference/stops', {
				errorMessage,
				isUnauthorized,
			});
			// #endregion
			res.statusCode = isUnauthorized ? 404 : 500;
			res.setHeader('Content-Type', 'application/json');
			res.end(JSON.stringify({
				error: isUnauthorized ? 'not authorized' : errorMessage,
			}));
		}
		return;
	}

	res.statusCode = 404;
	// #region agent log
	emitDebugLog('H5', 'ts-wrapper/src/api/server.ts:404', 'No matching route for request', {
		method: req.method || '',
		url: req.url,
	});
	// #endregion
	res.end('Not Found');
});

server.listen(port, () => {
	console.log(`Reference API running at http://localhost:${port}`);
}).on('error', (err: NodeJS.ErrnoException) => {
	if (err.code === 'EADDRINUSE') {
		console.error(`Port ${port} is already in use. Please set PORT environment variable to a different port.`);
		process.exit(1);
	}
	throw err;
});

process.on('SIGINT', async () => {
	await mongoClient?.close();
	process.exit(0);
});
