import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { environment } from "../environments/environment";
import { File } from './file';
import { People } from './people/people';

@Injectable({
  providedIn: 'root'
})
export class FilesService {
  constructor(private http: HttpClient) { }
  files$(dir: string) {
    return this.http.get<File[]>(`${environment.url}/api/files?dir=${dir}`)
  }
  peoples$() {
    return this.http.get<People[]>(`${environment.url}/api/peoples`)
  }
}
