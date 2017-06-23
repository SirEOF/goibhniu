import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-fermtrack',
  templateUrl: './fermtrack.component.html',
  styleUrls: ['./fermtrack.component.css']
})
export class FermtrackComponent {
	selectedID: number;

  	constructor() {
  		this.selectedID = -1;
  	}

  	handleListUpdate(event) {
    	this.selectedID = event;
  	}
}
