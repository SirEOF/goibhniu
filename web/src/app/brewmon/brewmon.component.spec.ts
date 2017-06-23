import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BrewmonComponent } from './brewmon.component';

describe('BrewmonComponent', () => {
  let component: BrewmonComponent;
  let fixture: ComponentFixture<BrewmonComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BrewmonComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BrewmonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
