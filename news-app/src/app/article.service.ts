import { Injectable } from '@angular/core';
import { Article } from './types/article';
import { Observable } from 'rxjs';
import { environment } from '../environments/environment';
import { CacheService } from './cache.service';

@Injectable({
  providedIn: 'root'
})
export class ArticleService {
  private _baseUrl: string;
  private _articleStorageKey: string = "articles";
  private _authorStorageKey: string = "authors";

  constructor(private _cache: CacheService) {
    this._baseUrl = environment.apiBaseUrl;
  }

  public getAllArticles(): Observable<Article[]> {
    const url = this._buildUrl("articles");

    return this._cache.get<Article[]>(url, this._articleStorageKey);
  }

  public getAllAuthors(): Observable<string[]> {
    const url = this._buildUrl("authors");

    return this._cache.get<string[]>(url, this._authorStorageKey);
  }

  private _buildUrl(endpoint: string): string {
    return this._baseUrl + endpoint;
  }

}
