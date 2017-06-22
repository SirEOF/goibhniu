import { Injectable } from '@angular/core';

import { Vessel } from './model/vessel';
import { VESSELS } from './model/vessel.mock';

import { Observable } from 'rxjs/Rx';

@Injectable()
export class VesselsService {
	constructor() { }

	getAll(): Observable<Vessel[]> {
		return Observable.of(VESSELS)
	}

	get(id: string): Observable<Vessel> {
		for (var i = 0, len = VESSELS.length; i < len; i++) {
			if(id === VESSELS[i].id) {
				return Observable.of(VESSELS[i]);
			}
		}
		return Observable.throw("id invalid");
	}
}
