import { Injectable } from '@angular/core';
import { CacheService } from './cache.service';
import { environment } from '../../environments/environment';
import { Article } from '../types/article';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ArticleService {
  private _baseUrl: string;
  private _storageKey: string = 'news-feed:articles';

  constructor(private _cache: CacheService) {
    this._baseUrl = environment.apiBaseUrl;
  }

  public getAll(): Observable<Article[]> {
    const url = this._baseUrl + 'articles';

    return this._cache.get<Article[]>(url, this._storageKey);
  }
}
