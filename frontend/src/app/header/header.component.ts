import { Component } from '@angular/core';


@Component({
  selector: 'app-header',
  template: `
    <div class="navbar is-white">
      <!-- logo -->
      <div class="navbar-brand">
        <a class ="navbar-item">
          <img src="assets/img/gator-logo.png">
      </a>
    </div>

    <!-- MENU STUFF -->
    <div class= "navbar-menu">
    <div class= "navbar-end">
      <a class= "navbar-item" routerLink=""> Home </a>
      <a class= "navbar-item" routerLink="Signup"> Signup </a>
      <a class= "navbar-item" routerLink="SellerPage"> I want to Sell! </a>
    </div>
    </div>
    
  <div class="container content has-text-centered">
       <p> WELCOME TO SWAMPY SELLS </p>
    </div>
  </div>
  <nav class="breadcrumb is-centered" aria-label="breadcrumbs">
  <ul>
    <li>
    <div class="dropdown is-hoverable">
  <div class="dropdown-trigger">
    <button class="button" aria-haspopup="true" aria-controls="dropdown-menu4">
      <span>Textbooks</span>
      <span class="icon is-small">
        <i class="fas fa-angle-down" aria-hidden="true"></i>
      </span>
    </button>
  </div>
  <div class="dropdown-menu" id="dropdown-menu4" role="menu">
    <div class="dropdown-content">
      <div class="dropdown-item">
        <p> Hardcover Textbooks</p>
      </div>
      <div class="dropdown-item">
        <p> Planners</p>
      </div>
      <div class="dropdown-item">
        <p> Novels</p>
      </div>
    </div>
  </div>
</div>
    </li>
    <li> <div class="dropdown is-hoverable">
  <div class="dropdown-trigger">
    <button class="button" aria-haspopup="true" aria-controls="dropdown-menu4">
      <span>Technology</span>
      <span class="icon is-small">
        <i class="fas fa-angle-down" aria-hidden="true"></i>
      </span>
    </button>
  </div>
  <div class="dropdown-menu" id="dropdown-menu4" role="menu">
    <div class="dropdown-content">
      <div class="dropdown-item">
        <p> Computers</p>
      </div>
      <div class="dropdown-item">
        <p> Phones</p>
      </div>
      <div class="dropdown-item">
        <p> Accessories</p>
      </div>
      <div class="dropdown-item">
        <p> IPads</p>
      </div>
      <div class="dropdown-item">
        <p> Headphones</p>
      </div>
    </div>
  </div>
</div>
</li>
    <li>
      <div class="dropdown is-hoverable">
  <div class="dropdown-trigger">
    <button class="button" aria-haspopup="true" aria-controls="dropdown-menu4">
      <span>Clothes</span>
      <span class="icon is-small">
        <i class="fas fa-angle-down" aria-hidden="true"></i>
      </span>
    </button>
  </div>
  <div class="dropdown-menu" id="dropdown-menu4" role="menu">
    <div class="dropdown-content">
      <div class="dropdown-item">
        <p> Women's Clothing</p>
      </div>
      <div class="dropdown-item">
        <p> Men's Clothing</p>
      </div>
      <div class="dropdown-item">
        <p> Shoes</p>
      </div>
      <div class="dropdown-item">
        <p> Accessories</p>
      </div>
    </div>
  </div>
</div>
</li>
    <li>
      <div class="dropdown is-hoverable">
  <div class="dropdown-trigger">
    <button class="button" aria-haspopup="true" aria-controls="dropdown-menu4">
      <span>General Decor</span>
      <span class="icon is-small">
        <i class="fas fa-angle-down" aria-hidden="true"></i>
      </span>
    </button>
  </div>
  <div class="dropdown-menu" id="dropdown-menu4" role="menu">
    <div class="dropdown-content">
      <div class="dropdown-item">
        <p> Posters</p>
      </div>
      <div class="dropdown-item">
        <p> Neon Lights</p>
      </div>
      <div class="dropdown-item">
        <p> Mirrors</p>
      </div>
      <div class="dropdown-item">
        <p> Decorative Pillows</p>
      </div>
      <div class="dropdown-item">
        <p> Blankers</p>
      </div>
    </div>
  </div>
</div>
</li>
    <li><div class="dropdown is-hoverable">
  <div class="dropdown-trigger">
    <button class="button" aria-haspopup="true" aria-controls="dropdown-menu4">
      <span>Furniture</span>
      <span class="icon is-small">
        <i class="fas fa-angle-down" aria-hidden="true"></i>
      </span>
    </button>
  </div>
  <div class="dropdown-menu" id="dropdown-menu4" role="menu">
    <div class="dropdown-content">
      <div class="dropdown-item">
        <p> Couches</p>
      </div>
      <div class="dropdown-item">
        <p> Barstools</p>
      </div>
      <div class="dropdown-item">
        <p> Desks</p>
      </div>
      <div class="dropdown-item">
        <p> Bedframes</p>
      </div>
      <div class="dropdown-item">
        <p> Chairs</p>
      </div>
      <div class="dropdown-item">
        <p> Outdoor Furniture</p>
      </div>
    </div>
  </div>
</div>
</li>
    <li>
      <div class="dropdown is-hoverable">
  <div class="dropdown-trigger">
    <button class="button" aria-haspopup="true" aria-controls="dropdown-menu4">
      <span>Tickets</span>
      <span class="icon is-small">
        <i class="fas fa-angle-down" aria-hidden="true"></i>
      </span>
    </button>
  </div>
  <div class="dropdown-menu" id="dropdown-menu4" role="menu">
    <div class="dropdown-content">
      <div class="dropdown-item">
        <p> Sports Tickets</p>
      </div>
      <div class="dropdown-item">
        <p> Theatre Tickets</p>
      </div>
      <div class="dropdown-item">
        <p> Random</p>
      </div>
    </div>
  </div>
</div>
</li>
<li>
      <div class="dropdown is-hoverable">
  <div class="dropdown-trigger">
    <button class="button" aria-haspopup="true" aria-controls="dropdown-menu4">
      <span>School Supplies</span>
      <span class="icon is-small">
        <i class="fas fa-angle-down" aria-hidden="true"></i>
      </span>
    </button>
  </div>
  <div class="dropdown-menu" id="dropdown-menu4" role="menu">
    <div class="dropdown-content">
      <div class="dropdown-item">
        <p> Notebooks</p>
      </div>
      <div class="dropdown-item">
        <p> Study Edge Packets</p>
      </div>
      <div class="dropdown-item">
        <p> Smokin'Notes Packets</p>
      </div>
      <div class="dropdown-item">
        <p> Writing Utensils</p>
      </div>
      <div class="dropdown-item">
        <p> Miscellaneous</p>
      </div>
    </div>
  </div>
</div>
</li>
  </ul>
</nav>
<div class="buttons is-link is-left">

<nav>
  <ul>
  <li><button [routerLink]="['/Login']">Log In</button></li>
  </ul>
  </nav>
</div>
`,
  styles: [`
    .container{
      font-family: "cool";
      font-size: 45px;
      
  }
  .buttons {
  border: none;
  color: blue;
  padding: 10px 12px;
  text-align: left;
  text-decoration: none;
  display: left;
  font-size: 16px;
}
  `]

})
export class HeaderComponent {

}
