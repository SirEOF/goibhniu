import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FermentableDetailComponent } from './fermentable-detail.component';

describe('FermentableDetailComponent', () => {
  let component: FermentableDetailComponent;
  let fixture: ComponentFixture<FermentableDetailComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FermentableDetailComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FermentableDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
