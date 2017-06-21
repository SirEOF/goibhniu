import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FermtrackComponent } from './fermtrack.component';

describe('FermtrackComponent', () => {
  let component: FermtrackComponent;
  let fixture: ComponentFixture<FermtrackComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FermtrackComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FermtrackComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
