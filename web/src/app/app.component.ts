import { Component } from '@angular/core';
import { Router, NavigationStart } from '@angular/router';
import { NavbarService } from './navbar/navbar.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [
  	NavbarService
  ]
})

export class AppComponent {
	constructor(public nav: NavbarService, private _router: Router) {
	}	

	ngOnInit() {
		this._router.events
			.subscribe((event) => {
				if (event instanceof NavigationStart) {
					if (event.url === "/") {
						this.nav.hide()
					} else {
						this.nav.show()
					}
				}
			})
	}
}
