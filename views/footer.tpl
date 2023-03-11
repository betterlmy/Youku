  <footer>
    <div class="author">
      Demo Official website:
      <a href="https://{{.Website}}">{{.Website}}</a> /
      Contact me:
      {{if eq .IsEmail 1}}

      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
      {{else}}
        <a class="email">禁止显示</a>
      {{end}}
    </div>
  </footer>