import { Injectable } from '@angular/core';

import { Fermentable } from './model/fermentable';
import { FERMENTABLES } from './model/fermentable.mock';

import { Observable } from 'rxjs/Rx';

@Injectable()
export class FermentablesService {
	constructor() { }

	getAll(): Observable<Fermentable[]> {
		return Observable.of(FERMENTABLES)
	}

	get(id: number): Observable<Fermentable> {
		for (var i = 0, len = FERMENTABLES.length; i < len; i++) {
			if(id === FERMENTABLES[i].id) {
				return Observable.of(FERMENTABLES[i]);
			}
		}
		return Observable.throw("id invalid");
	}
}
