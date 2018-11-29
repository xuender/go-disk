import { Observable } from 'rxjs';
import { Component, OnInit } from '@angular/core';
import { environment } from "../../environments/environment";

import { FilesService } from '../files.service';
import { People } from './people';

@Component({
  selector: 'app-people',
  templateUrl: './people.component.html',
  styleUrls: ['./people.component.scss']
})
export class PeopleComponent implements OnInit {
  peoples$: Observable<People[]>
  constructor(private filesService: FilesService) {
    this.peoples$ = filesService.peoples$()
  }

  url(id: string, index: number) {
    return `${environment.url}/api/peoples/${index}?id=${encodeURIComponent(id)}`
  }

  ngOnInit() {
  }

}
