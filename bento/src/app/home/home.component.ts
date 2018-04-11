import { Component, OnInit } from '@angular/core';
import { AmaiService } from '../amai.service';
import { Products } from '../Products';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
    /*we created a class named Products as a kind of template that will
     *contain the same info as the JSON from the API returns
     */
    items : Products[];

    constructor(private service : AmaiService ) {
    }

    ngOnInit() {
      this.service.getProductsObservable().subscribe(res => {
        this.items = res;
      });
    }

}
