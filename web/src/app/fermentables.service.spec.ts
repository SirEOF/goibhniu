import { TestBed, inject } from '@angular/core/testing';

import { FermentablesService } from './fermentables.service';

describe('FermentablesService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [FermentablesService]
    });
  });

  it('should be created', inject([FermentablesService], (service: FermentablesService) => {
    expect(service).toBeTruthy();
  }));
});
