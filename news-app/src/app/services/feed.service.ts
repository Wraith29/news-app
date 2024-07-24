import { Injectable } from '@angular/core';
import { CacheService } from './cache.service';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';
import { Feed } from '../types/feed';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class FeedService {
  private _baseUrl: string;
  private _storageKey: string = 'news-feed:feeds';

  constructor(private _cache: CacheService, private _http: HttpClient) {
    this._baseUrl = environment.apiBaseUrl;
  }

  public getAll(): Observable<Feed[]> {
    const url = this._baseUrl + 'feeds';

    return this._cache.get<Feed[]>(url, this._storageKey);
  }

  public delete(id: number): void {
    const url = this._baseUrl + 'feed';

    this._http.delete(url, {
      params: { feedId: id }
    }).subscribe({
      next: (value: Object) => {
        console.log(`Received: ${JSON.stringify(value)} from delete`);
      },
      error: (err: Error) => {
        console.error(`Received ${JSON.stringify(err)} from delete`);
      }
    });
  }

  public update(feed: Feed): void {
    const url = this._baseUrl + 'feed';

    this._http.put(url, {
      body: JSON.stringify(feed),
    }).subscribe({
      next: (value: Object) => {
        console.log(`Received: ${JSON.stringify(value)} from update`);
      },
      error: (err: Error) => {
        console.error(`Received ${JSON.stringify(err)} from update`);
      }
    })
  }
}
