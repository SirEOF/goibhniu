import { TestBed, inject } from '@angular/core/testing';

import { VesselsService } from './vessels.service';

describe('VesselsService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [VesselsService]
    });
  });

  it('should be created', inject([VesselsService], (service: VesselsService) => {
    expect(service).toBeTruthy();
  }));
});
