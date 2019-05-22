import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LoaningListComponent } from './loaning-list.component';

describe('LoaningListComponent', () => {
  let component: LoaningListComponent;
  let fixture: ComponentFixture<LoaningListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LoaningListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LoaningListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
