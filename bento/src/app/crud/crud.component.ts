import { Component, OnInit } from '@angular/core';
import { AmaiService } from '../amai.service';
import { CreateProductComponent } from '../create-product/create-product.component';
import { User } from '../User';

@Component({
  selector: 'app-crud',
  templateUrl: './crud.component.html',
  styleUrls: ['./crud.component.css'],
})
export class CrudComponent implements OnInit {
    /*we created a class named Products as a kind of template that will
     *contain the same info as the JSON from the API returns
     */
    users : User[];

    constructor(private service : AmaiService ) {
    }

    onSubmit( id: number ){
        console.log("on submit");
        this.service.deleteProduct(id);
        this.update();
    }

    refresh($event){
        this.update();
    }

    update(){
        this.service.getProductsObservable().subscribe(res => {
            this.users = res;
      });
    }

    ngOnInit() {
        this.update();
    }

}
