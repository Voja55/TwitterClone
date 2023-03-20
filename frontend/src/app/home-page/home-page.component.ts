import { Component} from '@angular/core';
import { StoreService } from '../services/store-service.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent {

  constructor(public store : StoreService) { }

}
