import booleanValid from '@turf/boolean-valid';
import { point } from '@turf/helpers';

/**
 * Validates stop coordinates using Turf.js.
 * Matches the logic in stop_lat_validation.go and stop_lon_validation.go:
 * - Latitude: -90 to 90
 * - Longitude: -180 to 180
 *
 * Uses GeoJSON Point + booleanValid to ensure coordinates conform to OGC Simple Feature Specification.
 *
 * @param lon - Longitude (-180 to 180)
 * @param lat - Latitude (-90 to 90)
 * @returns true if coordinates are valid
 */
export function isValidStopCoordinate(lon: number, lat: number): boolean {
	if (!Number.isFinite(lon) || !Number.isFinite(lat)) {
		return false;
	}
	try {
		const pt = point([lon, lat]);
		return booleanValid(pt);
	} catch {
		return false;
	}
}
