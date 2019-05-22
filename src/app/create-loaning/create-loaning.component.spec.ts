import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateLoaningComponent } from './create-loaning.component';

describe('CreateLoaningComponent', () => {
  let component: CreateLoaningComponent;
  let fixture: ComponentFixture<CreateLoaningComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CreateLoaningComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateLoaningComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
