import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <!-- header -->
    <app-header></app-header>

    <app-home></app-home>
    <app-profiles></app-profiles>
    <!-- routes get injected here -->
    
    <router-outlet></router-outlet>

    <!-- footer -->

    <app-footer></app-footer>

  `,
  styles: []
})
export class AppComponent {
  title = 'my-angular-site';
}
