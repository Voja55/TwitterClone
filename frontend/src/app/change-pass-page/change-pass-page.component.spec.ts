import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChangePassPageComponent } from './change-pass-page.component';

describe('ChangePassPageComponent', () => {
  let component: ChangePassPageComponent;
  let fixture: ComponentFixture<ChangePassPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ChangePassPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ChangePassPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
