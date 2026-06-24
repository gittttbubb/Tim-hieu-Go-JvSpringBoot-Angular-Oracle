import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PermissionBadgeComponent } from './permission-badge.component';

describe('PermissionBadgeComponent', () => {
  let component: PermissionBadgeComponent;
  let fixture: ComponentFixture<PermissionBadgeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PermissionBadgeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PermissionBadgeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
