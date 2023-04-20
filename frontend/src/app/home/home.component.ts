import { Component } from '@angular/core';

@Component({
  selector: 'app-home',
  template: `
    <section class="hero is-primary is-bold is-fullheight">
  <div class="hero-body">
  
    
  </div>
<app-profiles></app-profiles>
</section>
  `,
  styles: [`
     .hero  {
        background-image: url('/assets/images/background2.jpg') !important;
        background-size: cover;
        background-position: center center;
     }
     `]

})

export class HomeComponent {

}
