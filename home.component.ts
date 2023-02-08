import { Component } from '@angular/core';

@Component({
  selector: 'app-home',
  template: `
    <section class="hero is-primary is-bold is-fullheight">
  <div class="hero-body">
  
    
  </div>
</section>
  `,
  styles: [`
     .hero  {
        background-image: url('/assets/img/background2.jpg') !important;
        background-size: cover;
        background-position: center center;
     }
     `]
})

export class HomeComponent {

}
