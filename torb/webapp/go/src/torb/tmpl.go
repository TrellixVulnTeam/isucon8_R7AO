package main

func indexTmpl(e, u, o []byte) [][]byte {
	return [][]byte{[]byte(`<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta charset="utf-8">
    <title>Torb</title>

    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <link rel="shortcut icon" href="`), o, []byte(`/favicon.ico" type="image/vnd.microsoft.icon" />
    <link rel="stylesheet" href="`), o, []byte(`/css/bootstrap.min.css">
    <link rel="stylesheet" href="`), o, []byte(`/css/layout.css">
  </head>
  <body>
    <div id="container" class="container">
      <div id="app-wrapper" data-login-user="`), u, []byte(`" data-events="`), e, []byte(`">

        <div id="menu-bar" class="d-flex flex-column flex-md-row align-items-center p-3 px-md-4 mb-3 bg-white border-bottom box-shadow">
          <h1 class="my-0 mr-md-auto font-weight-normal h5">Torb</h1>
          <nav class="my-2 my-md-0 mr-md-3">
            <a class="p-2 text-dark" href="#" v-on:click.stop.prevent="signIn" v-if="!currentUser">サインイン</a>
            <a class="p-2 text-dark" href="#" v-on:click.stop.prevent="signUp" v-if="!currentUser">サインアップ</a>
            <a class="p-2 text-dark" href="#" v-on:click.stop.prevent="showMyPage" v-if="currentUser">マイページ: {{ currentUser.nickname }}</a>
            <a class="p-2 text-dark" href="#" v-on:click.stop.prevent="signOut" v-if="currentUser">サインアウト</a>
          </nav>
        </div>

        <div class="events">
          <h3>開催中のイベント</h3>

          <div class="list-group">
            <a href="#" v-for="event in events" v-on:click.stop.prevent="open(event.id)" class="list-group-item" >
              <div class="d-flex w-100 justify-content-between">
                <h5 class="mb-1">{{ event.title }}</h5>
                <small class="text-muted">{{ event.remains }} / {{ event.total }}</small>
              </div>
              <span class="badge badge-dark" v-for="rank in ranks">{{ rank }} <small>{{ event.sheets[rank].price }}円</small></span>
            </a>
          </div>
        </div>

        <div class="modals">
          <div class="modal" id="confirm-modal" tabindex="-1" role="dialog">
            <div class="modal-dialog modal-sm" role="document">
              <div class="modal-content">
                <div class="modal-header">
                  <h5 class="modal-title">{{ title }}</h5>
                  <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                  </button>
                </div>
                <div class="modal-body">
                  <p>{{ message }}</p>
                </div>
                <div class="modal-footer">
                  <button type="button" class="btn btn-primary" v-on:click.stop.prevent="ok()" tabindex="0" role="button">OK</button>
                  <button type="button" class="btn btn-secondary" data-dismiss="modal">NG</button>
                </div>
              </div>
            </div>
          </div>

          <div class="modal" id="login-modal" tabindex="-1" role="dialog">
            <div class="modal-dialog" role="document">
              <form action="#" v-on:submit.stop.prevent="submit">
                <div class="modal-content">
                  <div class="modal-header">
                    <h5 class="modal-title">サインイン</h5>
                    <button type="button" class="close" data-dismiss="modal" aria-label="Cancel">
                      <span aria-hidden="true">&times;</span>
                    </button>
                  </div>
                  <div class="modal-body">
                    <div class="form-group">
                      <label for="login-form-login-name">ログインID</label>
                      <input type="text" class="form-control" id="login-form-login-name" placeholder="your_login_name" v-model="loginName" required>
                    </div>
                    <div class="form-group">
                      <label for="login-form-password">パスワード</label>
                      <input type="password" class="form-control" id="login-form-password" placeholder="********" v-model="password" required>
                    </div>
                  </div>
                  <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">キャンセル</button>
                    <button type="submit" class="btn btn-primary">サインイン</button>
                  </div>
                </div>
              </form>
            </div>
          </div>

          <div class="modal" id="register-modal" tabindex="-1" role="dialog">
            <div class="modal-dialog" role="document">
              <form action="#" v-on:submit.stop.prevent="submit">
                <div class="modal-content">
                  <div class="modal-header">
                    <h5 class="modal-title">サインアップ</h5>
                    <button type="button" class="close" data-dismiss="modal" aria-label="Cancel">
                      <span aria-hidden="true">&times;</span>
                    </button>
                  </div>
                  <div class="modal-body">
                    <div class="form-group">
                      <label for="register-form-nickname">名前</label>
                      <input type="text" class="form-control" id="register-form-nickname" placeholder="Your Name" v-model="nickname" required>
                    </div>
                    <div class="form-group">
                      <label for="register-form-login-name">ログインID</label>
                      <input type="text" class="form-control" id="register-form-login-name" placeholder="your_login_name" v-model="loginName" required>
                    </div>
                    <div class="form-group">
                      <label for="register-form-password">パスワード</label>
                      <input type="password" class="form-control" id="register-form-password" placeholder="********" v-model="password" required>
                    </div>
                  </div>
                  <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">キャンセル</button>
                    <button type="submit" class="btn btn-primary">登録</button>
                  </div>
                </div>
              </form>
            </div>
          </div>

          <div class="modal" id="event-modal" tabindex="-1" role="dialog">
            <div class="modal-dialog modal-lg event-modal-dialog" role="document">
              <div class="modal-content">
                <div class="modal-header">
                  <h5 class="modal-title">{{ event.title }}</h5>
                  <button type="button" class="close" data-dismiss="modal" aria-label="Cancel">
                    <span aria-hidden="true">&times;</span>
                  </button>
                </div>
                <div class="modal-body">
                  <div class="d-flex w-100">
                    <small class="text-muted">{{ event.remains }} / {{ event.total }}</small>
                  </div>
                  <div class="d-flex w-100" v-for="rank in ranks">
                    <span class="rank">{{ rank }}</span>
                    <div class="progress remaining-sheets-bar">
                      <div class="progress-bar" role="progressbar" v-bind:aria-valuenow="event.sheets[rank].remains" aria-valuemin="0" v-bind:aria-valuemax="event.sheets[rank].total" v-bind:style="{ width: 100 * (event.sheets[rank].remains/event.sheets[rank].total) + '%' }">{{ event.sheets[rank].remains }}</div>
                    </div>
                  </div>
                  <div class="sheets-tables">
                    <table class="table" v-for="rank in ranks">
                      <thead>
                        <tr>
                          <th colspan="25">{{ rank }}</td>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="n in divRange(event.sheets[rank].total, 25)">
                          <td v-for="i in 25" v-bind:class="{ 'table-dark': event.sheets[rank].detail[(n-1)*25+(i-1)].reserved }">
                            <span v-bind:class="{ 'mine': event.sheets[rank].detail[(n-1)*25+(i-1)].mine }" v-on:click.stop.prevent="freeSheet(rank, event.sheets[rank].detail[(n-1)*25+(i-1)].num)">
                              {{ event.sheets[rank].detail[(n-1)*25+(i-1)].num }}
                            </span>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
                <div class="modal-footer">
                  <div class="btn-group" role="group" aria-label="Reserve sheet">
                    <button type="buttom" class="btn btn-primary" v-for="rank in ranks" v-bind:disabled="isSoldOut(rank)" v-on:click.stop.prevent="reserveSheet(rank)">{{ rank }}席 {{ event.sheets[rank].price }}円</button>
                  </div>
                  <button type="button" class="btn btn-secondary" data-dismiss="modal">閉じる</button>
                </div>
              </div>
            </div>
          </div>

          <div class="modal" id="my-page-modal" tabindex="-1" role="dialog">
            <div class="modal-dialog modal-lg my-page-modal-dialog" role="document">
              <div class="modal-content">
                <div class="modal-header">
                  <h5 class="modal-title">マイページ: {{ user.nickname }}</h5>
                  <button type="button" class="close" data-dismiss="modal" aria-label="Cancel">
                    <span aria-hidden="true">&times;</span>
                  </button>
                </div>
                <div class="modal-body">
                  <div class="row">
                    <dt class="col-sm-2">予約総額</dt>
                    <dd class="col-sm-10">{{ user.total_price }}円</dd>
                  </div>
                  <div class="row">
                    <div class="col">
                      <h5 class="modal-title">最近予約したイベント</h5>
                      <div class="list-group">
                        <a href="#" v-for="event in user.recent_events" v-on:click.stop.prevent="openEvent(event)" class="list-group-item" >
                          <div>
                            <h5 class="mb-1">{{ event.title }}</h5>
                            <small class="text-muted">{{ event.remains }} / {{ event.total }} (<span v-text="event.closed ? '終了' : event.public ? '公開中' : '非公開'"></span>）</small>
                          </div>
                          <span class="badge badge-dark" v-for="rank in ranks">{{ rank }} <small>{{ event.sheets[rank].price }}円</small></span>
                        </a>
                        <div class="d-flex w-100 justify-content-between" v-if="user.recent_events.length === 0">
                          まだ予約済のイベントはありません
                        </div>
                      </div>
                    </div>
                    <div class="col">
                      <h5 class="modal-title">最近予約した席</h5>
                      <div class="list-group">
                        <a href="#" v-for="reservation in user.recent_reservations" v-on:click.stop.prevent="openEvent(reservation.event)" class="list-group-item" >
                          <div>
                            <h5 class="mb-1">{{ reservation.event.title }}</h5>
                            <small class="text-muted"><span v-text="reservation.event.closed ? '終了' : reservation.event.public ? '公開中' : '非公開'"></span></small>
                          </div>
                          <small class="text-muted">{{ formatDateTime(reservation.reserved_at) }}: {{ reservation.sheet_rank }}-{{ reservation.sheet_num }}<span v-if="reservation.canceled_at"><br />(キャンセル済: {{ formatDateTime(reservation.canceled_at) }}）</span></small>
                        </a>
                        <div class="d-flex w-100 justify-content-between" v-if="user.recent_reservations.length === 0">
                          まだ予約済の席はありません
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="modal-footer">
                  <button type="button" class="btn btn-secondary" data-dismiss="modal">閉じる</button>
                </div>
              </div>
            </div>
          </div>

        </div>

      </div><!-- /app-wrapper -->
    </div>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script type="text/javascript" src="`), o, []byte(`/js/jquery-3.3.1.slim.min.js"></script>
    <script type="text/javascript" src="`), o, []byte(`/js/bootstrap.bundle.min.js"></script>
    <script type="text/javascript" src="`), o, []byte(`/js/bootstrap-waitingfor.min.js"></script>
    <script type="text/javascript" src="`), o, []byte(`/js/vue.min.js"></script>
    <script type="text/javascript" src="`), o, []byte(`/js/fetch.min.js"></script>
    <script type="text/javascript" src="`), o, []byte(`/js/app.js"></script>
  </body>
</html>`)}
}

func adminTmpl(e, u, o []byte) [][]byte {
	return [][]byte{[]byte(`<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta charset="utf-8">
    <title>Torb管理</title>

    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <link rel="shortcut icon" href="`), o, []byte(`/favicon.ico" type="image/vnd.microsoft.icon" />
    <link rel="stylesheet" href="`), o, []byte(`/css/bootstrap.min.css">
    <link rel="stylesheet" href="`), o, []byte(`/css/admin.css">
  </head>
  <body>
    <div id="container" class="container">

      <div id="app-wrapper" data-administrator="`), u, []byte(`" data-events="`), e, []byte(`">

        <nav id="menu-bar" class="navbar navbar-expand-lg navbar-light bg-light">
          <h1 class="navbar-brand h5">Torb管理</h1>
          <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#header-navbar-content" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
          </button>

          <div id="header-navbar-content" class="collapse navbar-collapse">
            <ul class="navbar-nav mr-auto">
              <li class="nav-item" v-if="currentAdministrator">
                <a class="p-2 text-dark" href="#" v-on:click.stop.prevent="downloadSalesReport">総合購買レポート</a>
              </li>
            </ul>
            <div class="form-inline my-2 my-lg-0">
              <a class="p-2 text-dark" href="#" v-on:click.stop.prevent="signIn" v-if="!currentAdministrator">サインイン</a>
              <a class="p-2 text-dark" href="#" v-on:click.stop.prevent="signOut" v-if="currentAdministrator">サインアウト ({{ currentAdministrator.nickname }})</a>
            </div>
          </div>
        </div>

        <div class="events" v-if="isAdmin">
          <h3>すべてのイベント</h3>

          <div class="list-group">
            <a href="#" v-for="event in events" v-on:click.stop.prevent="open(event.id)" class="list-group-item">
              <div class="d-flex w-100 justify-content-between">
                <h5 class="mb-1">{{ event.title }}</h5>
                <small class="text-muted">{{ event.remains }} / {{ event.total }} (<span v-text="event.closed ? '終了' : event.public ? '公開中' : '非公開'"></span>）</small>
              </div>
              <span class="badge badge-dark" v-for="rank in ranks">{{ rank }} <small>{{ event.sheets[rank].price }}円</small></span>
            </a>
          </div>

          <div class="events-actions">
            <button type="button" class="btn btn-primary btn-lg btn-block" v-on:click.stop.prevent="openEventRegistrationModal">新規イベント</button>
          </div>
        </div>

        <div class="modals">
          <div class="modal" id="confirm-modal" tabindex="-1" role="dialog">
            <div class="modal-dialog modal-sm" role="document">
              <div class="modal-content">
                <div class="modal-header">
                  <h5 class="modal-title">{{ title }}</h5>
                  <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                  </button>
                </div>
                <div class="modal-body">
                  <p>{{ message }}</p>
                </div>
                <div class="modal-footer">
                  <button type="button" class="btn btn-primary" v-on:click.stop.prevent="ok()" tabindex="0" role="button">OK</button>
                  <button type="button" class="btn btn-secondary" data-dismiss="modal">NG</button>
                </div>
              </div>
            </div>
          </div>

          <div class="modal" id="login-modal" tabindex="-1" role="dialog">
            <div class="modal-dialog" role="document">
              <form action="#" v-on:submit.stop.prevent="submit">
                <div class="modal-content">
                  <div class="modal-header">
                    <h5 class="modal-title">サインイン</h5>
                    <button type="button" class="close" data-dismiss="modal" aria-label="Cancel">
                      <span aria-hidden="true">&times;</span>
                    </button>
                  </div>
                  <div class="modal-body">
                    <div class="form-group">
                      <label for="login-form-login-name">ログインID</label>
                      <input type="text" class="form-control" id="login-form-login-name" placeholder="your_login_name" v-model="loginName" required>
                    </div>
                    <div class="form-group">
                      <label for="login-form-password">パスワード</label>
                      <input type="password" class="form-control" id="login-form-password" placeholder="********" v-model="password" required>
                    </div>
                  </div>
                  <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">キャンセル</button>
                    <button type="submit" class="btn btn-primary">サインイン</button>
                  </div>
                </div>
              </form>
            </div>
          </div>

          <div class="modal" id="event-registration-modal" tabindex="-1" role="dialog">
            <div class="modal-dialog" role="document">
              <form action="#" v-on:submit.stop.prevent="submit">
                <div class="modal-content">
                  <div class="modal-header">
                    <h5 class="modal-title">新規イベント</h5>
                    <button type="button" class="close" data-dismiss="modal" aria-label="Cancel">
                      <span aria-hidden="true">&times;</span>
                    </button>
                  </div>
                  <div class="modal-body">
                    <div class="form-group">
                      <label for="event-registration-form-title">タイトル</label>
                      <input type="text" class="form-control" id="event-registration-form-title" placeholder="The Catcher in the Rye" name="title" v-model="title" required>
                    </div>
                    <div class="form-group">
                      <label for="event-registration-form-price">ベース価格</label>
                      <input type="number" class="form-control" id="event-registration-form-price" placeholder="1000" name="price" v-model="price" min="1000" max="30000" required>
                    </div>
                    <div class="form-group">
                      <div class="form-check">
                        <input class="form-check-input" type="radio" name="public" v-model="public" id="event-registration-form-public" value="1" required>
                        <label class="form-check-label" for="event-registration-form-public">公開</label>
                      </div>
                      <div class="form-check">
                        <input class="form-check-input" type="radio" name="public" v-model="public" id="event-registration-form-private" value="0" required>
                        <label class="form-check-label" for="event-registration-form-private">非公開</label>
                      </div>
                    </div>
                  </div>
                  <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">キャンセル</button>
                    <button type="submit" class="btn btn-primary">作成</button>
                  </div>
                </div>
              </form>
            </div>
          </div>

          <div class="modal" id="event-modal" tabindex="-1" role="dialog">
            <div class="modal-dialog modal-lg event-modal-content" role="document">
              <div class="modal-content">
                <div class="modal-header">
                  <h5 class="modal-title">{{ event.title }}</h5>
                  <button type="button" class="close" data-dismiss="modal" aria-label="Cancel">
                    <span aria-hidden="true">&times;</span>
                  </button>
                </div>
                <div class="modal-body">
                  <div class="d-flex w-100">
                    <small class="text-muted">{{ event.remains }} / {{ event.total }} (<span v-text="event.closed ? '終了' : event.public ? '公開中' : '非公開'"></span>）</small>
                  </div>
                  <div class="d-flex w-100" v-for="rank in ranks">
                    <span class="rank">{{ rank }}</span>
                    <div class="progress remaining-sheets-bar">
                      <div class="progress-bar" role="progressbar" v-bind:aria-valuenow="event.sheets[rank].remains" aria-valuemin="0" v-bind:aria-valuemax="event.sheets[rank].total" v-bind:style="{ width: 100 * (event.sheets[rank].remains/event.sheets[rank].total) + '%' }">{{ event.sheets[rank].remains }}</div>
                    </div>
                  </div>
                  <div class="sheets-tables">
                    <table class="table" v-for="rank in ranks">
                      <thead>
                        <tr>
                          <th colspan="25">{{ rank }}</td>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="n in divRange(event.sheets[rank].total, 25)">
                          <td v-for="i in 25" v-bind:class="{ 'table-dark': event.sheets[rank].detail[(n-1)*25+(i-1)].reserved }">
                            {{ event.sheets[rank].detail[(n-1)*25+(i-1)].num }}
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
                <div class="modal-footer">
                  <button type="buttom" class="btn btn-info" v-on:click.stop.prevent="downloadSalesReport">購買レポート</button>
                  <button type="buttom" class="btn btn-warning" v-if="!event.closed && !event.public" v-on:click.stop.prevent="publish">公開する</button>
                  <button type="buttom" class="btn btn-danger" v-if="!event.closed && !event.public" v-on:click.stop.prevent="close">終了する</button>
                  <button type="buttom" class="btn btn-danger" v-if="event.public" v-on:click.stop.prevent="disappear">公開を停止する</button>
                  <button type="button" class="btn btn-secondary" data-dismiss="modal">閉じる</button>
                </div>
              </div>
            </div>
          </div>
        </div>

      </div><!-- /app-wrapper -->
    </div>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script type="text/javascript" src="`), o, []byte(`/js/jquery-3.3.1.slim.min.js"></script>
    <script type="text/javascript" src="`), o, []byte(`/js/bootstrap.bundle.min.js"></script>
    <script type="text/javascript" src="`), o, []byte(`/js/bootstrap-waitingfor.min.js"></script>
    <script type="text/javascript" src="`), o, []byte(`/js/vue.min.js"></script>
    <script type="text/javascript" src="`), o, []byte(`/js/fetch.min.js"></script>
    <script type="text/javascript" src="`), o, []byte(`/js/admin.js"></script>
  </body>
</html>`)}
}
