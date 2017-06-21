import { BrowserModule } from '@angular/platform-browser';
import { Routes, RouterModule } from '@angular/router';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { FermentablesListComponent } from './fermentables-list/fermentables-list.component';
import { FermentableDetailComponent } from './fermentable-detail/fermentable-detail.component';
import { FermtrackComponent } from './fermtrack/fermtrack.component';
import { HomeComponent } from './home/home.component';
import { NavbarComponent } from './navbar/navbar.component';

export const ROUTES: Routes = [
	{ path: '', component: HomeComponent },
	{ path: 'ferm', component: FermtrackComponent},
	{ path: 'ferm/:id', component: FermtrackComponent }
];

@NgModule({
  declarations: [
    AppComponent,
    FermentablesListComponent,
    FermentableDetailComponent,
    FermtrackComponent,
    HomeComponent,
    NavbarComponent
  ],
  imports: [
    BrowserModule,
    NgbModule.forRoot(),
    RouterModule.forRoot(ROUTES)
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
