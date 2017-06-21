import { Reading } from './reading';

export class Fermentable {
	id: number;
	recipe: string;
	tags: string[];
	fermenter: string;
	startdate: string;
	enddate: string;
	readings: Reading[];
	notes: string;
}