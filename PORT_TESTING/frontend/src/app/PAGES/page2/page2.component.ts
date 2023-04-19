import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';


@Component({
  selector: 'app-page2',
  templateUrl: './page2.component.html',
  styleUrls: ['./page2.component.css']
})

export class Page2Component implements OnInit{
  public msg: any
  constructor( private http: HttpClient ) { }
  public ngOnInit(): void {
    this.msg = this.http.get("/page2").subscribe((response)=>{
      this.msg = response;
    })
  }
}

