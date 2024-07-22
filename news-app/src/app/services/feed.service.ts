import { Injectable } from '@angular/core';
import { CacheService } from './cache.service';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';
import { Feed } from '../types/feed';

@Injectable({
  providedIn: 'root',
})
export class FeedService {
  private _baseUrl: string;
  private _storageKey: string = 'news-feed:feeds';

  constructor(private _cache: CacheService) {
    this._baseUrl = environment.apiBaseUrl;
  }

  public getAll(): Observable<Feed[]> {
    const url = this._baseUrl + 'feeds';

    return this._cache.get<Feed[]>(url, this._storageKey);
  }
}
