import { Component, Input, OnInit, OnChanges, SimpleChanges } from '@angular/core';
import { DatePipe } from '@angular/common';

import { Chart } from 'angular-highcharts';

import { Fermentable } from '../model/fermentable';
import { FermentablesService } from '../fermentables.service';
import { Vessel } from '../model/vessel';
import { VesselsService } from '../vessels.service';
import { Reading } from '../model/reading';


@Component({
  selector: 'app-fermentable-detail',
  providers: [
  	FermentablesService, 
  	VesselsService,
  	DatePipe
  ],
  templateUrl: './fermentable-detail.component.html',
  styleUrls: ['./fermentable-detail.component.css']
})

export class FermentableDetailComponent implements OnChanges, OnInit {

  @Input() selected: number;
  brew: Fermentable;
  vessel: Vessel
  error: string;
  hide: boolean;
  fermChart: Chart;

  constructor(private _fermentablesService: FermentablesService, 
  			  private _vesselsService: VesselsService,
  			  private _datePipe: DatePipe) { 
  	this.hide = true;
  }

  updateChart(readings: Reading[]) {
  	var series: number[][] = [];

  	for(var i = 0; i < readings.length; i++) {
  		series.push([Date.parse(readings[i].updatetime), readings[i].gravity])
  	}

  	this.fermChart = new Chart({
		chart: {
    	    type: 'line'
    	},
		title: {
    	    text: 'Specific Gravity'
		},
		legend: {
			enabled: false
		},
		xAxis: {
			type: 'datetime',
			labels: {
				format: '{value:%b %e, %Y}'
			},
		},
		plotOptions: {
			line: {
				dataLabels: {
					enabled: true
				},
			}
		},
		series: [
			{
				data: series
			}
		]
	});
  }

  getVesselInfo(id: string) {
  	if(id !== undefined) {
	  	this._vesselsService.get(id).subscribe(
	  		value => this.vessel = value,
	  	)
  	}
  }

  getFermentationDetailValue(value: Fermentable) {
  	this.brew = value;
  	this.getVesselInfo(this.brew.fermenter);
  	this.updateChart(this.brew.readings);
  }

  getFermentationDetail(id: number) {
  	if(id >= 0) {
		this._fermentablesService.get(id).subscribe(
			value => this.getFermentationDetailValue(value), 
		)
		this.hide = false;
	} else {
		this.hide = true;
	}
  }

  ngOnChanges(changes: SimpleChanges) {
  	if (changes.selected && !changes.selected.isFirstChange()) {
  		console.log("onChange", changes.selected.currentValue )
  		this.getFermentationDetail(changes.selected.currentValue);
  	}
  }

  ngOnInit() {
  	if(this.selected !== undefined) {
  		this.getFermentationDetail(this.selected);
  	}
  }
}
