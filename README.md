## Film Libriary  
Just execute **make build** then **make run** and enjoy(If u dont need some specific option for building a compose).  
  
If there is no need in testing then just exclude from makefile init-dep block installing of mockery.  
  
If u want to use Admin endpoinds u have to change user role directly in database and then log in again. Command u should use:  
**UPDATE users SET role_id = 1 WHERE id = '{UUID of the user}';**  

Router is in /internal/app/app.go

.env file is in the root of project.  
  
⚡ Сделано без мам, пап и copilot'ов  
