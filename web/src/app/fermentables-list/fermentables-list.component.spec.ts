import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FermentablesListComponent } from './fermentables-list.component';

describe('FermentablesListComponent', () => {
  let component: FermentablesListComponent;
  let fixture: ComponentFixture<FermentablesListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FermentablesListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FermentablesListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
