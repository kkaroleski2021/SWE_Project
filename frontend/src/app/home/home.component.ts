import { Component } from '@angular/core';

@Component({
  selector: 'a-component',
  template: `
    <section class="hero is-primary is-bold is-fullheight">
  <div class="hero-body">
      <b-component></b-component>
    
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

export class AComponent {}

