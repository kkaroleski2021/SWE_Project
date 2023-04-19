import { Component } from '@angular/core';

@Component({
  selector: 'd-component',
  template: `
   <section class="hero is-primary is-bold is-fullheight">
  <div class="hero-body">
    <div class = "container">
      <c-component></c-component>

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

export class DComponent {}

