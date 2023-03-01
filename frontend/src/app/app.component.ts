import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <!-- header -->
    <app-header></app-header>
    <a-component></a-component>
    
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
