# Git Rule

個人開発においては以下の限りではない

## Branch

- 本番環境 : `main`
- 開発環境 : `develop`
- 機能追加 : `feature/*`
- バグ修正 : `fix/*`
- 機能改善 : `refactor/*`

## Push

- main ブランチを取り込むときは merge commit を残さないために `fetch → rebase` を実行する 
- レビュー後 merge するときは commit を意味のある単位にまとめて `rebase → squash` を実行する  
- `rebase` 後は `push --force-with-lease` を実行する
