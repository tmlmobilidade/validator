import { dirname, join } from 'path';
import { fileURLToPath } from 'url';

import { GTFSValidator } from '../index.js';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const result = await GTFSValidator(join(__dirname, '../../data/Bom.zip'));
console.log(result);
