import { Component } from '@angular/core';

@Component({
  selector: 'a-component',
  template: `
   <section class="hero is-primary is-bold is-fullheight">
  <div class="hero-body">
    <div class = "container">
      <b-component></b-component>

  </div>
  </div>
</section>


  `,
  styles: [`
  .hero  {
        background-image: url('/assets/background2.jpg') !important;
        background-size: cover;
        background-position: center center;
     }
     `]
})

export class AComponent {}

