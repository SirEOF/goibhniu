import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { Observable } from 'rxjs/Rx';

import { Fermentable } from '../model/fermentable';
import { FermentablesService } from '../fermentables.service';

@Component({
  selector: 'app-fermentables-list',
  providers: [FermentablesService],
  templateUrl: './fermentables-list.component.html',
  styleUrls: ['./fermentables-list.component.css']
})
export class FermentablesListComponent implements OnInit {
	brews: Observable<Fermentable[]>;

  @Output() change: EventEmitter<number> = new EventEmitter<number>();

  constructor(private _fermentablesService: FermentablesService) { }

  getListing() {
		this.brews = this._fermentablesService.getAll()
  }

  ngOnInit() {
		this.getListing()
  }

  handleClick(id: number) {
    this.change.emit(id)
  }
}
