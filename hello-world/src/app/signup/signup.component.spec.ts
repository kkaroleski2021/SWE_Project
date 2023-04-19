import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SComponent } from './signup.component';

describe('SignupComponent', () => {
  let component: SComponent;
  let fixture: ComponentFixture<SComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
