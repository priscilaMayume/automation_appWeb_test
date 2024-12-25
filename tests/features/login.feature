Funcionalidade: Login
  Esquema do Cenário: Fazer login
    Quando eu faço login com "<email>" e "<password>"
    Então eu recebo status <status>

    Exemplos:
      | email               | password | status |
      | eve.holt@reqres.in  | pistol   | 200    |
      | user@notfound.in    | 12345    | 400    |
