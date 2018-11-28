import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { environment } from "../environments/environment";
import { File } from './file';

@Injectable({
  providedIn: 'root'
})
export class FilesService {
  constructor(private http: HttpClient) { }
  files$(dir: string) {
    return this.http.get<File[]>(`${environment.url}/api/files?dir=${dir}`)
  }
}
