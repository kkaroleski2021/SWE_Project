// import { Component } from '@angular/core';

// @Component({
//   selector: 'app-root',
//   templateUrl: './app.component.html',
//   styleUrls: ['./app.component.css']
// })
// export class AppComponent {
//   title = 'the software engineering project of all time';
// }

import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <!-- header -->
    <app-header></app-header>

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