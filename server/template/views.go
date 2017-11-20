// Code generated by go generate; DO NOT EDIT.
// 2017-11-20 15:11:09.531826622 -0800 PST m=+0.008960222

package template

var templateViewsMap = map[string]string{
	"about": `{{ define "title"}}{{ t "About" }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "About" }}</h1>
    <ul>
        <li>
            <a href="{{ route "settings" }}">{{ t "Settings" }}</a>
        </li>
        <li>
            <a href="{{ route "sessions" }}">{{ t "Sessions" }}</a>
        </li>
        {{ if .user.IsAdmin }}
        <li>
            <a href="{{ route "users" }}">{{ t "Users" }}</a>
        </li>
        {{ end }}
    </ul>
</section>

<div class="panel">
    <h3>{{ t "Version" }}</h3>
    <ul>
        <li><strong>{{ t "Version:" }}</strong> {{ .version }}</li>
        <li><strong>{{ t "Build Date:" }}</strong> {{ .build_date }}</li>
    </ul>
</div>

<div class="panel">
    <h3>{{ t "Authors" }}</h3>
    <ul>
        <li><strong>{{ t "Author:" }}</strong> Frédéric Guillot</li>
        <li><strong>{{ t "License:" }}</strong> Apache 2.0</li>
    </ul>
</div>

{{ end }}
`,
	"add_subscription": `{{ define "title"}}{{ t "New Subscription" }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "New Subscription" }}</h1>
    <ul>
        <li>
            <a href="{{ route "feeds" }}">{{ t "Feeds" }}</a>
        </li>
        <li>
            <a href="{{ route "export" }}">{{ t "Export" }}</a>
        </li>
        <li>
            <a href="{{ route "import" }}">{{ t "Import" }}</a>
        </li>
    </ul>
</section>

{{ if not .categories }}
    <p class="alert alert-error">{{ t "There is no category. You must have at least one category." }}</p>
{{ else }}
    <form action="{{ route "submitSubscription" }}" method="post" autocomplete="off">
        <input type="hidden" name="csrf" value="{{ .csrf }}">

        {{ if .errorMessage }}
            <div class="alert alert-error">{{ t .errorMessage }}</div>
        {{ end }}

        <label for="form-url">{{ t "URL" }}</label>
        <input type="url" name="url" id="form-url" placeholder="https://domain.tld/" value="{{ .form.URL }}" required autofocus>

        <label for="form-category">{{ t "Category" }}</label>
        <select id="form-category" name="category_id">
            {{ range .categories }}
                <option value="{{ .ID }}">{{ .Title }}</option>
            {{ end }}
        </select>

        <div class="buttons">
            <button type="submit" class="button button-primary" data-label-loading="{{ t "Loading..." }}">{{ t "Find a subscription" }}</button>
        </div>
    </form>
{{ end }}

{{ end }}
`,
	"categories": `{{ define "title"}}{{ t "Categories" }} ({{ .total }}){{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "Categories" }} ({{ .total }})</h1>
    <ul>
        <li>
            <a href="{{ route "createCategory" }}">{{ t "Create a category" }}</a>
        </li>
    </ul>
</section>

{{ if not .categories }}
    <p class="alert alert-error">{{ t "There is no category." }}</p>
{{ else }}
    <div class="items">
        {{ range .categories }}
        <article class="item">
            <div class="item-header">
                <span class="item-title">
                    <a href="{{ route "categoryEntries" "categoryID" .ID }}">{{ .Title }}</a>
                </span>
            </div>
            <div class="item-meta">
                <ul>
                    <li>
                        {{ if eq .FeedCount 0 }}
                            {{ t "No feed." }}
                        {{ else }}
                            {{ plural "plural.categories.feed_count" .FeedCount .FeedCount }}
                        {{ end }}
                    </li>
                </ul>
                <ul>
                    <li>
                        <a href="{{ route "editCategory" "categoryID" .ID }}">{{ t "Edit" }}</a>
                    </li>
                    {{ if eq .FeedCount 0 }}
                    <li>
                        <a href="{{ route "removeCategory" "categoryID" .ID }}">{{ t "Remove" }}</a>
                    </li>
                    {{ end }}
                </ul>
            </div>
        </article>
        {{ end }}
    </div>
{{ end }}

{{ end }}
`,
	"category_entries": `{{ define "title"}}{{ .category.Title }} ({{ .total }}){{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ .category.Title }} ({{ .total }})</h1>
    <ul>
        <li>
            <a href="#" data-on-click="markPageAsRead">{{ t "Mark this page as read" }}</a>
        </li>
    </ul>
</section>

{{ if not .entries }}
    <p class="alert">{{ t "There is no article in this category." }}</p>
{{ else }}
    <div class="items">
        {{ range .entries }}
        <article class="item item-status-{{ .Status }}" data-id="{{ .ID }}">
            <div class="item-header">
                <span class="item-title">
                    {{ if ne .Feed.Icon.IconID 0 }}
                        <img src="{{ route "icon" "iconID" .Feed.Icon.IconID }}" width="16" height="16">
                    {{ end }}
                    <a href="{{ route "categoryEntry" "categoryID" .Feed.Category.ID "entryID" .ID }}">{{ .Title }}</a>
                </span>
                <span class="category"><a href="{{ route "categoryEntries" "categoryID" .Feed.Category.ID }}">{{ .Feed.Category.Title }}</a></span>
            </div>
            <div class="item-meta">
                <ul>
                    <li>
                        <a href="{{ route "feedEntries" "feedID" .Feed.ID }}" title="{{ .Feed.Title }}">{{ domain .Feed.SiteURL }}</a>
                    </li>
                    <li>
                        <time datetime="{{ isodate .Date }}" title="{{ isodate .Date }}">{{ elapsed .Date }}</time>
                    </li>
                    <li>
                        <a href="{{ .URL }}" target="_blank" rel="noopener noreferrer" referrerpolicy="no-referrer" data-original-link="true">{{ t "Original" }}</a>
                    </li>
                </ul>
            </div>
        </article>
        {{ end }}
    </div>
    {{ template "pagination" .pagination }}
{{ end }}

{{ end }}
`,
	"choose_subscription": `{{ define "title"}}{{ t "Choose a Subscription" }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "New Subscription" }}</h1>
    <ul>
        <li>
            <a href="{{ route "feeds" }}">{{ t "Feeds" }}</a>
        </li>
        <li>
            <a href="{{ route "export" }}">{{ t "Export" }}</a>
        </li>
        <li>
            <a href="{{ route "import" }}">{{ t "Import" }}</a>
        </li>
    </ul>
</section>

<form action="{{ route "chooseSubscription" }}" method="POST">
    <input type="hidden" name="csrf" value="{{ .csrf }}">
    <input type="hidden" name="category_id" value="{{ .categoryID }}">

    <h3>{{ t "Choose a Subscription" }}</h3>

    {{ range .subscriptions }}
        <div class="radio-group">
            <label title="{{ .URL }}"><input type="radio" name="url" value="{{ .URL }}"> {{ .Title }}</label> ({{ .Type }})
            <small title="Type = {{ .Type }}"><a href="{{ .URL }}" target="_blank" rel="noopener noreferrer" referrerpolicy="no-referrer">{{ .URL }}</a></small>
        </div>
    {{ end }}

    <div class="buttons">
        <button type="submit" class="button button-primary" data-label-loading="{{ t "Loading..." }}">{{ t "Subscribe" }}</button>
    </div>
</form>
{{ end }}
`,
	"create_category": `{{ define "title"}}{{ t "New Category" }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "New Category" }}</h1>
    <ul>
        <li>
            <a href="{{ route "categories" }}">{{ t "Categories" }}</a>
        </li>
    </ul>
</section>

<form action="{{ route "saveCategory" }}" method="post" autocomplete="off">
    <input type="hidden" name="csrf" value="{{ .csrf }}">

    {{ if .errorMessage }}
        <div class="alert alert-error">{{ t .errorMessage }}</div>
    {{ end }}

    <label for="form-title">{{ t "Title" }}</label>
    <input type="text" name="title" id="form-title" value="{{ .form.Title }}" required autofocus>

    <div class="buttons">
        <button type="submit" class="button button-primary" data-label-loading="{{ t "Loading..." }}">{{ t "Save" }}</button> {{ t "or" }} <a href="{{ route "categories" }}">{{ t "cancel" }}</a>
    </div>
</form>
{{ end }}
`,
	"create_user": `{{ define "title"}}{{ t "New User" }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "New User" }}</h1>
    <ul>
        <li>
            <a href="{{ route "settings" }}">{{ t "Settings" }}</a>
        </li>
        <li>
            <a href="{{ route "sessions" }}">{{ t "Sessions" }}</a>
        </li>
        <li>
            <a href="{{ route "users" }}">{{ t "Users" }}</a>
        </li>
    </ul>
</section>

<form action="{{ route "saveUser" }}" method="post" autocomplete="off">
    <input type="hidden" name="csrf" value="{{ .csrf }}">

    {{ if .errorMessage }}
        <div class="alert alert-error">{{ t .errorMessage }}</div>
    {{ end }}

    <label for="form-username">{{ t "Username" }}</label>
    <input type="text" name="username" id="form-username" value="{{ .form.Username }}" required autofocus>

    <label for="form-password">{{ t "Password" }}</label>
    <input type="password" name="password" id="form-password" value="{{ .form.Password }}" required>

    <label for="form-confirmation">{{ t "Confirmation" }}</label>
    <input type="password" name="confirmation" id="form-confirmation" value="{{ .form.Confirmation }}" required>

    <label><input type="checkbox" name="is_admin" value="1" {{ if .form.IsAdmin }}checked="checked"{{ end }}> {{ t "Administrator" }}</label>

    <div class="buttons">
        <button type="submit" class="button button-primary" data-label-loading="{{ t "Loading..." }}">{{ t "Save" }}</button> {{ t "or" }} <a href="{{ route "users" }}">{{ t "cancel" }}</a>
    </div>
</form>
{{ end }}
`,
	"edit_category": `{{ define "title"}}{{ t "Edit Category: %s" .category.Title }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "Edit Category: %s" .category.Title }}</h1>
    <ul>
        <li>
            <a href="{{ route "categories" }}">{{ t "Categories" }}</a>
        </li>
        <li>
            <a href="{{ route "createCategory" }}">{{ t "Create a category" }}</a>
        </li>
    </ul>
</section>

<form action="{{ route "updateCategory" "categoryID" .category.ID }}" method="post" autocomplete="off">
    <input type="hidden" name="csrf" value="{{ .csrf }}">

    {{ if .errorMessage }}
        <div class="alert alert-error">{{ t .errorMessage }}</div>
    {{ end }}

    <label for="form-title">{{ t "Title" }}</label>
    <input type="text" name="title" id="form-title" value="{{ .form.Title }}" required autofocus>

    <div class="buttons">
        <button type="submit" class="button button-primary" data-label-loading="{{ t "Loading..." }}">{{ t "Update" }}</button> {{ t "or" }} <a href="{{ route "categories" }}">{{ t "cancel" }}</a>
    </div>
</form>
{{ end }}
`,
	"edit_feed": `{{ define "title"}}{{ t "Edit Feed: %s" .feed.Title }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ .feed.Title }}</h1>
    <ul>
        <li>
            <a href="{{ route "feeds" }}">{{ t "Feeds" }}</a>
        </li>
        <li>
            <a href="{{ route "addSubscription" }}">{{ t "Add subscription" }}</a>
        </li>
        <li>
            <a href="{{ route "export" }}">{{ t "Export" }}</a>
        </li>
        <li>
            <a href="{{ route "import" }}">{{ t "Import" }}</a>
        </li>
    </ul>
</section>

{{ if not .categories }}
    <p class="alert alert-error">{{ t "There is no category!" }}</p>
{{ else }}
    {{ if ne .feed.ParsingErrorCount 0 }}
    <div class="alert alert-error">
        <h3>{{ t "Last Parsing Error" }}</h3>
        {{ .feed.ParsingErrorMsg }}
    </div>
    {{ end }}

    <form action="{{ route "updateFeed" "feedID" .feed.ID }}" method="post" autocomplete="off">
        <input type="hidden" name="csrf" value="{{ .csrf }}">

        {{ if .errorMessage }}
            <div class="alert alert-error">{{ t .errorMessage }}</div>
        {{ end }}

        <label for="form-title">{{ t "Title" }}</label>
        <input type="text" name="title" id="form-title" value="{{ .form.Title }}" required autofocus>

        <label for="form-site-url">{{ t "Site URL" }}</label>
        <input type="url" name="site_url" id="form-site-url" placeholder="https://domain.tld/" value="{{ .form.SiteURL }}" required>

        <label for="form-feed-url">{{ t "Feed URL" }}</label>
        <input type="url" name="feed_url" id="form-feed-url" placeholder="https://domain.tld/" value="{{ .form.FeedURL }}" required>

        <label for="form-category">{{ t "Category" }}</label>
        <select id="form-category" name="category_id">
        {{ range .categories }}
            <option value="{{ .ID }}" {{ if eq .ID $.form.CategoryID }}selected="selected"{{ end }}>{{ .Title }}</option>
        {{ end }}
        </select>

        <div class="buttons">
            <button type="submit" class="button button-primary" data-label-loading="{{ t "Loading..." }}">{{ t "Update" }}</button> {{ t "or" }} <a href="{{ route "feeds" }}">{{ t "cancel" }}</a>
        </div>
    </form>
{{ end }}

{{ end }}`,
	"edit_user": `{{ define "title"}}{{ t "Edit user: %s" .selected_user.Username }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "Edit user %s" .selected_user.Username }}"</h1>
    <ul>
        <li>
            <a href="{{ route "settings" }}">{{ t "Settings" }}</a>
        </li>
        <li>
            <a href="{{ route "sessions" }}">{{ t "Sessions" }}</a>
        </li>
        <li>
            <a href="{{ route "users" }}">{{ t "Users" }}</a>
        </li>
        <li>
            <a href="{{ route "createUser" }}">{{ t "Add user" }}</a>
        </li>
    </ul>
</section>

<form action="{{ route "updateUser" "userID" .selected_user.ID }}" method="post" autocomplete="off">
    <input type="hidden" name="csrf" value="{{ .csrf }}">

    {{ if .errorMessage }}
        <div class="alert alert-error">{{ t .errorMessage }}</div>
    {{ end }}

    <label for="form-username">{{ t "Username" }}</label>
    <input type="text" name="username" id="form-username" value="{{ .form.Username }}" required autofocus>

    <label for="form-password">{{ t "Password" }}</label>
    <input type="password" name="password" id="form-password" value="{{ .form.Password }}">

    <label for="form-confirmation">{{ t "Confirmation" }}</label>
    <input type="password" name="confirmation" id="form-confirmation" value="{{ .form.Confirmation }}">

    <label><input type="checkbox" name="is_admin" value="1" {{ if .form.IsAdmin }}checked="checked"{{ end }}> {{ t "Administrator" }}</label>

    <div class="buttons">
        <button type="submit" class="button button-primary" data-label-loading="{{ t "Loading..." }}">{{ t "Update" }}</button> {{ t "or" }} <a href="{{ route "users" }}">{{ t "cancel" }}</a>
    </div>
</form>
{{ end }}
`,
	"entry": `{{ define "title"}}{{ .entry.Title }}{{ end }}

{{ define "content"}}
<section class="entry">
    <header class="entry-header">
        <h1>
            <a href="{{ .entry.URL }}" target="_blank" rel="noopener noreferrer" referrerpolicy="no-referrer">{{ .entry.Title }}</a>
        </h1>
        <div class="entry-meta">
            <span class="entry-website">
                {{ if ne .entry.Feed.Icon.IconID 0 }}
                    <img src="{{ route "icon" "iconID" .entry.Feed.Icon.IconID }}" width="16" height="16">
                {{ end }}
                <a href="{{ route "feedEntries" "feedID" .entry.Feed.ID }}">{{ .entry.Feed.Title }}</a>
            </span>
            {{ if .entry.Author }}
                <span class="entry-author">
                    {{ if contains .entry.Author "@" }}
                        - <a href="mailto:{{ .entry.Author }}">{{ .entry.Author }}</a>
                    {{ else }}
                        – <em>{{ .entry.Author }}</em>
                    {{ end }}
                </span>
            {{ end }}
            <span class="category">
                <a href="{{ route "categoryEntries" "categoryID" .entry.Feed.Category.ID }}">{{ .entry.Feed.Category.Title }}</a>
            </span>
        </div>
        <div class="entry-date">
            <time datetime="{{ isodate .entry.Date }}" title="{{ isodate .entry.Date }}">{{ elapsed .entry.Date }}</time>
        </div>
    </header>
    <div class="pagination-top">
        {{ template "entry_pagination" . }}
    </div>
    <article class="entry-content">
        {{ noescape (proxyFilter .entry.Content) }}
    </article>
    {{ if .entry.Enclosures }}
    <aside class="entry-enclosures">
        <h3>{{ t "Attachments" }}</h3>
        {{ range .entry.Enclosures }}
            <div class="entry-enclosure">
                {{ if hasPrefix .MimeType "audio/" }}
                    <div class="enclosure-audio">
                        <audio controls preload="metadata">
                            <source src="{{ .URL }}" type="{{ .MimeType }}">
                        </audio>
                    </div>
                {{ else if hasPrefix .MimeType "video/" }}
                    <div class="enclosure-video">
                        <video controls preload="metadata">
                            <source src="{{ .URL }}" type="{{ .MimeType }}">
                        </video>
                    </div>
                {{ else if hasPrefix .MimeType "image/" }}
                    <div class="enclosure-image">
                        <img src="{{ .URL }}" title="{{ .URL }} ({{ .MimeType }})" alt="{{ .URL }} ({{ .MimeType }})">
                    </div>
                {{ end }}

                <div class="entry-enclosure-download">
                    <a href="{{ .URL }}" title="{{ .URL }} ({{ .MimeType }})" target="_blank" rel="noopener noreferrer" referrerpolicy="no-referrer">{{ t "Download" }}</a>
                    <small>({{ .URL }})</small>
                </div>
            </div>
        {{ end }}
    </aside>
    {{ end }}
</section>

<div class="pagination-bottom">
    {{ template "entry_pagination" . }}
</div>
{{ end }}
`,
	"feed_entries": `{{ define "title"}}{{ .feed.Title }} ({{ .total }}){{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ .feed.Title }} ({{ .total }})</h1>
    <ul>
        <li>
            <a href="{{ route "refreshFeed" "feedID" .feed.ID }}">{{ t "Refresh" }}</a>
        </li>
        <li>
            <a href="{{ route "editFeed" "feedID" .feed.ID }}">{{ t "Edit" }}</a>
        </li>
        <li>
            <a href="#" data-on-click="markPageAsRead">{{ t "Mark this page as read" }}</a>
        </li>
    </ul>
</section>

{{ if ne .feed.ParsingErrorCount 0 }}
<div class="alert alert-error">
    <h3>{{ t "There is a problem with this feed" }}</h3>
    {{ .feed.ParsingErrorMsg }}
</div>
{{ else if not .entries }}
    <p class="alert">{{ t "There is no article for this feed." }}</p>
{{ else }}
    <div class="items">
        {{ range .entries }}
        <article class="item item-status-{{ .Status }}" data-id="{{ .ID }}">
            <div class="item-header">
                <span class="item-title">
                    {{ if ne .Feed.Icon.IconID 0 }}
                        <img src="{{ route "icon" "iconID" .Feed.Icon.IconID }}" width="16" height="16">
                    {{ end }}
                    <a href="{{ route "feedEntry" "feedID" .Feed.ID "entryID" .ID }}">{{ .Title }}</a>
                </span>
                <span class="category"><a href="{{ route "categoryEntries" "categoryID" .Feed.Category.ID }}">{{ .Feed.Category.Title }}</a></span>
            </div>
            <div class="item-meta">
                <ul>
                    <li>
                        <a href="{{ route "feedEntries" "feedID" .Feed.ID }}" title="{{ .Feed.Title }}">{{ domain .Feed.SiteURL }}</a>
                    </li>
                    <li>
                        <time datetime="{{ isodate .Date }}" title="{{ isodate .Date }}">{{ elapsed .Date }}</time>
                    </li>
                    <li>
                        <a href="{{ .URL }}" target="_blank" rel="noopener noreferrer" referrerpolicy="no-referrer" data-original-link="true">{{ t "Original" }}</a>
                    </li>
                </ul>
            </div>
        </article>
        {{ end }}
    </div>
    {{ template "pagination" .pagination }}
{{ end }}

{{ end }}
`,
	"feeds": `{{ define "title"}}{{ t "Feeds" }} ({{ .total }}){{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "Feeds" }} ({{ .total }})</h1>
    <ul>
        <li>
            <a href="{{ route "addSubscription" }}">{{ t "Add subscription" }}</a>
        </li>
        <li>
            <a href="{{ route "export" }}">{{ t "Export" }}</a>
        </li>
        <li>
            <a href="{{ route "import" }}">{{ t "Import" }}</a>
        </li>
    </ul>
</section>

{{ if not .feeds }}
    <p class="alert">{{ t "You don't have any subscription." }}</p>
{{ else }}
    <div class="items">
        {{ range .feeds }}
        <article class="item">
            <div class="item-header">
                <span class="item-title">
                    {{ if ne .Icon.IconID 0 }}
                        <img src="{{ route "icon" "iconID" .Icon.IconID }}" width="16" height="16">
                    {{ end }}
                    <a href="{{ route "feedEntries" "feedID" .ID }}">{{ .Title }}</a>
                </span>
                <span class="category">
                    <a href="{{ route "categoryEntries" "categoryID" .Category.ID }}">{{ .Category.Title }}</a>
                </span>
            </div>
            <div class="item-meta">
                <ul>
                    <li>
                        <a href="{{ .SiteURL }}" target="_blank" rel="noopener noreferrer" referrerpolicy="no-referrer" data-original-link="true">{{ domain .SiteURL }}</a>
                    </li>
                    <li>
                        {{ t "Last check:" }} <time datetime="{{ isodate .CheckedAt }}" title="{{ isodate .CheckedAt }}">{{ elapsed .CheckedAt }}</time>
                    </li>
                    {{ if ne .ParsingErrorCount 0 }}
                        <li><strong title="{{ .ParsingErrorMsg }}">{{ plural "plural.feed.error_count" .ParsingErrorCount .ParsingErrorCount }}</strong></li>
                    {{ end }}
                </ul>
                <ul>
                    <li>
                        <a href="{{ route "refreshFeed" "feedID" .ID }}">{{ t "Refresh" }}</a>
                    </li>
                    <li>
                        <a href="{{ route "editFeed" "feedID" .ID }}">{{ t "Edit" }}</a>
                    </li>
                    <li>
                        <a href="{{ route "removeFeed" "feedID" .ID }}">{{ t "Remove" }}</a>
                    </li>
                </ul>
            </div>
        </article>
        {{ end }}
    </div>
{{ end }}

{{ end }}
`,
	"history": `{{ define "title"}}{{ t "History" }} ({{ .total }}){{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "History" }} ({{ .total }})</h1>
</section>

{{ if not .entries }}
    <p class="alert alert-info">{{ t "There is no history at the moment." }}</p>
{{ else }}
    <div class="items">
        {{ range .entries }}
        <article class="item item-status-{{ .Status }}" data-id="{{ .ID }}">
            <div class="item-header">
                <span class="item-title">
                    {{ if ne .Feed.Icon.IconID 0 }}
                        <img src="{{ route "icon" "iconID" .Feed.Icon.IconID }}" width="16" height="16">
                    {{ end }}
                    <a href="{{ route "readEntry" "entryID" .ID }}">{{ .Title }}</a>
                </span>
                <span class="category"><a href="{{ route "categoryEntries" "categoryID" .Feed.Category.ID }}">{{ .Feed.Category.Title }}</a></span>
            </div>
            <div class="item-meta">
                <ul>
                    <li>
                        <a href="{{ route "feedEntries" "feedID" .Feed.ID }}" title="{{ .Feed.Title }}">{{ domain .Feed.SiteURL }}</a>
                    </li>
                    <li>
                        <time datetime="{{ isodate .Date }}" title="{{ isodate .Date }}">{{ elapsed .Date }}</time>
                    </li>
                    <li>
                        <a href="{{ .URL }}" target="_blank" rel="noopener noreferrer" referrerpolicy="no-referrer" data-original-link="true">{{ t "Original" }}</a>
                    </li>
                </ul>
            </div>
        </article>
        {{ end }}
    </div>
    {{ template "pagination" .pagination }}
{{ end }}

{{ end }}
`,
	"import": `{{ define "title"}}{{ t "Import" }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "Import" }}</h1>
    <ul>
        <li>
            <a href="{{ route "feeds" }}">{{ t "Feeds" }}</a>
        </li>
        <li>
            <a href="{{ route "addSubscription" }}">{{ t "Add subscription" }}</a>
        </li>
        <li>
            <a href="{{ route "export" }}">{{ t "Export" }}</a>
        </li>
    </ul>
</section>

<form action="{{ route "uploadOPML" }}" method="post" enctype="multipart/form-data">
    <input type="hidden" name="csrf" value="{{ .csrf }}">

    {{ if .errorMessage }}
        <div class="alert alert-error">{{ t .errorMessage }}</div>
    {{ end }}

    <label for="form-file">{{ t "OPML file" }}</label>
    <input type="file" name="file" id="form-file">

    <div class="buttons">
        <button type="submit" class="button button-primary" data-label-loading="{{ t "Loading..." }}">{{ t "Import" }}</button>
    </div>
</form>

{{ end }}
`,
	"login": `{{ define "title"}}{{ t "Sign In" }}{{ end }}

{{ define "content"}}
<section class="login-form">
    <form action="{{ route "checkLogin" }}" method="post">
        <input type="hidden" name="csrf" value="{{ .csrf }}">

        {{ if .errorMessage }}
            <div class="alert alert-error">{{ t .errorMessage }}</div>
        {{ end }}

        <label for="form-username">{{ t "Username" }}</label>
        <input type="text" name="username" id="form-username" required autofocus>

        <label for="form-password">{{ t "Password" }}</label>
        <input type="password" name="password" id="form-password" required>

        <div class="buttons">
            <button type="submit" class="button button-primary" data-label-loading="{{ t "Loading..." }}">{{ t "Sign in" }}</button>
        </div>
    </form>
</section>
{{ end }}
`,
	"sessions": `{{ define "title"}}{{ t "Sessions" }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "Sessions" }}</h1>
    <ul>
        <li>
            <a href="{{ route "settings" }}">{{ t "Settings" }}</a>
        </li>
        <li>
            <a href="{{ route "users" }}">{{ t "Users" }}</a>
        </li>
        <li>
            <a href="{{ route "createUser" }}">{{ t "Add user" }}</a>
        </li>
    </ul>
</section>

<table class="table-overflow">
    <tr>
        <th>{{ t "Date" }}</th>
        <th>{{ t "IP Address" }}</th>
        <th>{{ t "User Agent" }}</th>
        <th>{{ t "Actions" }}</th>
    </tr>
    {{ range .sessions }}
    <tr {{ if eq .Token $.currentSessionToken }}class="row-highlighted"{{ end }}>
        <td class="column-20" title="{{ isodate .CreatedAt }}">{{ elapsed .CreatedAt }}</td>
        <td class="column-20" title="{{ .IP }}">{{ .IP }}</td>
        <td title="{{ .UserAgent }}">{{ .UserAgent }}</td>
        <td class="column-20">
            {{ if eq .Token $.currentSessionToken }}
                {{ t "Current session" }}
            {{ else }}
                <a href="{{ route "removeSession" "sessionID" .ID }}">{{ t "Remove" }}</a>
            {{ end }}
        </td>
    </tr>
    {{ end }}
</table>

{{ end }}
`,
	"settings": `{{ define "title"}}{{ t "Settings" }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "Settings" }}</h1>
    <ul>
        <li>
            <a href="{{ route "sessions" }}">{{ t "Sessions" }}</a>
        </li>
        {{ if .user.IsAdmin }}
        <li>
            <a href="{{ route "users" }}">{{ t "Users" }}</a>
        </li>
        {{ end }}
        <li>
            <a href="{{ route "about" }}">{{ t "About" }}</a>
        </li>
    </ul>
</section>

<form method="post" autocomplete="off" action="{{ route "updateSettings" }}">
    <input type="hidden" name="csrf" value="{{ .csrf }}">

    {{ if .errorMessage }}
        <div class="alert alert-error">{{ t .errorMessage }}</div>
    {{ end }}

    <label for="form-username">{{ t "Username" }}</label>
    <input type="text" name="username" id="form-username" value="{{ .form.Username }}" required>

    <label for="form-password">{{ t "Password" }}</label>
    <input type="password" name="password" id="form-password" value="{{ .form.Password }}" autocomplete="off">

    <label for="form-confirmation">{{ t "Confirmation" }}</label>
    <input type="password" name="confirmation" id="form-confirmation" value="{{ .form.Confirmation }}" autocomplete="off">

    <label for="form-language">{{ t "Language" }}</label>
    <select id="form-language" name="language">
    {{ range $key, $value := .languages }}
        <option value="{{ $key }}" {{ if eq $key $.form.Language }}selected="selected"{{ end }}>{{ $value }}</option>
    {{ end }}
    </select>

    <label for="form-timezone">{{ t "Timezone" }}</label>
    <select id="form-timezone" name="timezone">
    {{ range $key, $value := .timezones }}
        <option value="{{ $key }}" {{ if eq $key $.form.Timezone }}selected="selected"{{ end }}>{{ $value }}</option>
    {{ end }}
    </select>

    <label for="form-theme">{{ t "Theme" }}</label>
    <select id="form-theme" name="theme">
    {{ range $key, $value := .themes }}
        <option value="{{ $key }}" {{ if eq $key $.form.Theme }}selected="selected"{{ end }}>{{ $value }}</option>
    {{ end }}
    </select>

    <div class="buttons">
        <button type="submit" class="button button-primary" data-label-loading="{{ t "Loading..." }}">{{ t "Update" }}</button>
    </div>
</form>

{{ end }}
`,
	"unread": `{{ define "title"}}{{ t "Unread Items" }} {{ if gt .countUnread 0 }}({{ .countUnread }}){{ end }} {{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "Unread" }} ({{ .countUnread }})</h1>
    <ul>
        <li>
            <a href="#" data-on-click="markPageAsRead">{{ t "Mark this page as read" }}</a>
        </li>
    </ul>
</section>

{{ if not .entries }}
    <p class="alert">{{ t "There is no unread article." }}</p>
{{ else }}
    <div class="items hide-read-items">
        {{ range .entries }}
        <article class="item item-status-{{ .Status }}" data-id="{{ .ID }}">
            <div class="item-header">
                <span class="item-title">
                    {{ if ne .Feed.Icon.IconID 0 }}
                        <img src="{{ route "icon" "iconID" .Feed.Icon.IconID }}" width="16" height="16">
                    {{ end }}
                    <a href="{{ route "unreadEntry" "entryID" .ID }}">{{ .Title }}</a>
                </span>
                <span class="category"><a href="{{ route "categoryEntries" "categoryID" .Feed.Category.ID }}">{{ .Feed.Category.Title }}</a></span>
            </div>
            <div class="item-meta">
                <ul>
                    <li>
                        <a href="{{ route "feedEntries" "feedID" .Feed.ID }}" title="{{ .Feed.Title }}">{{ domain .Feed.SiteURL }}</a>
                    </li>
                    <li>
                        <time datetime="{{ isodate .Date }}" title="{{ isodate .Date }}">{{ elapsed .Date }}</time>
                    </li>
                    <li>
                        <a href="{{ .URL }}" target="_blank" rel="noopener noreferrer" referrerpolicy="no-referrer" data-original-link="true">{{ t "Original" }}</a>
                    </li>
                </ul>
            </div>
        </article>
        {{ end }}
    </div>
    {{ template "pagination" .pagination }}
{{ end }}

{{ end }}`,
	"users": `{{ define "title"}}{{ t "Users" }}{{ end }}

{{ define "content"}}
<section class="page-header">
    <h1>{{ t "Users" }}</h1>
    <ul>
        <li>
            <a href="{{ route "settings" }}">{{ t "Settings" }}</a>
        </li>
        <li>
            <a href="{{ route "sessions" }}">{{ t "Sessions" }}</a>
        </li>
        <li>
            <a href="{{ route "createUser" }}">{{ t "Add user" }}</a>
        </li>
    </ul>
</section>

{{ if eq (len .users) 1 }}
    <p class="alert">{{ t "You are the only user." }}</p>
{{ else }}
    <table>
        <tr>
            <th class="column-20">{{ t "Username" }}</th>
            <th>{{ t "Administrator" }}</th>
            <th>{{ t "Last Login" }}</th>
            <th>{{ t "Actions" }}</th>
        </tr>
        {{ range .users }}
            {{ if ne .ID $.user.ID }}
            <tr>
                <td>{{ .Username }}</td>
                <td>{{ if eq .IsAdmin true }}{{ t "Yes" }}{{ else }}{{ t "No" }}{{ end }}</td>
                <td>
                    {{ if .LastLoginAt }}
                        <time datetime="{{ isodate .LastLoginAt }}" title="{{ isodate .LastLoginAt }}">{{ elapsed .LastLoginAt }}</time>
                    {{ else }}
                        {{ t "Never" }}
                    {{ end }}
                </td>
                <td>
                    <a href="{{ route "editUser" "userID" .ID }}">{{ t "Edit" }}</a>,
                    <a href="{{ route "removeUser" "userID" .ID }}">{{ t "Remove" }}</a>
                </td>
            </tr>
            {{ end }}
        {{ end }}
    </table>
{{ end }}

{{ end }}
`,
}

var templateViewsMapChecksums = map[string]string{
	"about":               "56f1d45d8b9944306c66be0712320527e739a0ce4fccbd97a4c414c8f9cfab04",
	"add_subscription":    "098ea9e492e18242bd414b22c4d8638006d113f728e5ae78c9186663f60ae3f1",
	"categories":          "721b6bae6aa6461f4e020d667707fabe53c94b399f7d74febef2de5eb9f15071",
	"category_entries":    "0bdcf28ef29b976b78d1add431896a8c56791476abd7a4240998d52c3efe1f35",
	"choose_subscription": "d37682743d8bbd84738a964e238103db2651f95fa340c6e285ffe2e12548d673",
	"create_category":     "2b82af5d2dcd67898dc5daa57a6461e6ff8121a6089b2a2a1be909f35e4a2275",
	"create_user":         "966b31d0414e0d0a547ef9ada428cbd24a91100bfed491f780c0461892a2489b",
	"edit_category":       "cee720faadcec58289b707ad30af623d2ee66c1ce23a732965463250d7ff41c5",
	"edit_feed":           "c5bc4c22bf7e8348d880395250545595d21fb8c8e723fc5d7cca68e25d250884",
	"edit_user":           "f0f79704983de3ca7858bd8cda7a372c3999f5e4e0cf951fba5fa2c1752f9111",
	"entry":               "32e605edd6d43773ac31329d247ebd81d38d974cd43689d91de79fffec7fe04b",
	"feed_entries":        "9aff923b6c7452dec1514feada7e0d2bbc1ec21c6f5e9f48b2de41d1b731ffe4",
	"feeds":               "ddcf12a47c850e6a1f3b85c9ab6566b4e45adfcd7a3546381a0c3a7a54f2b7d4",
	"history":             "439000d0be8fd716f3b89860af4d721e05baef0c2ccd2325ba020c940d6aa847",
	"import":              "73b5112e20bfd232bf73334544186ea419505936bc237d481517a8622901878f",
	"login":               "568f2f69f248048f3e55e9bbc719077a74ae23fe18f237aa40e3de37e97b7a41",
	"sessions":            "7fcd3bb794d4ad01eb9fa515660f04c8e79e1568970fd541cc7b2de8a76e1542",
	"settings":            "9c89bfd70ff288b4256e5205be78a7645450b364db1df51d10fee3cb915b2c6b",
	"unread":              "b6f9be1a72188947c75a6fdcac6ff7878db7745f9efa46318e0433102892a722",
	"users":               "5bd535de3e46d9b14667d8159a5ec1478d6e028a77bf306c89d7b55813eeb625",
}
