import { Directive, Input, TemplateRef, ViewContainerRef, inject, } from '@angular/core';
import { AuthStore } from '../../store/auth.store';

@Directive({
    selector: '[appHasPermission]',
})
export class HasPermissionDirective {
    private templateRef = inject(TemplateRef<unknown>);
    private viewContainer = inject(ViewContainerRef);
    private authStore = inject(AuthStore);

    @Input() set appHasPermission(permission: | string | string[]) {
        this.viewContainer.clear();
        let allowed = false;
        if (typeof permission === 'string') {
            allowed = this.authStore.hasPermission(permission);
        } else {
            allowed = this.authStore.hasAnyPermission(permission);
        }
        if (allowed) {
            this.viewContainer.createEmbeddedView(this.templateRef);
        }
    }
}