import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterBusinessUserComponent } from './register-business-user.component';

describe('RegisterBusinessUserComponent', () => {
  let component: RegisterBusinessUserComponent;
  let fixture: ComponentFixture<RegisterBusinessUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegisterBusinessUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RegisterBusinessUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
