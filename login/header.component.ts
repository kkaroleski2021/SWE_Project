import { Component } from '@angular/core';


@Component({
  selector: 'app-header',
  template: `

    <div class="navbar is-white">
      <!-- logo -->
      <div class="navbar-brand">
    
    </div>


    
  <div class="container content has-text-centered">
       <p> Log In to Swampy Sells! </p>
    </div>
  </div>


`,
  styles: [`
    .container{
      font-family: "cool";
      font-size: 45px;
      
  }

  `]

})
export class HeaderComponent {

}
