import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserOverridesComponent } from './user-overrides.component';

describe('UserOverridesComponent', () => {
  let component: UserOverridesComponent;
  let fixture: ComponentFixture<UserOverridesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [UserOverridesComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UserOverridesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
