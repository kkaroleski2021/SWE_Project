import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-page1',
  templateUrl: './page1.component.html',
  styleUrls: ['./page1.component.css']
})

export class Page1Component implements OnInit{
  public msg: any
  constructor( private http: HttpClient ) { }
  public ngOnInit(): void {
    this.getMessage().subscribe(res=>{
      if(res){ 
        console.log("response")
      }
      console.log(res)
      this.msg = res;
    })
  }
  getMessage(){
    return this.http.get("http://localhost:9000/page1");
  }
}


