################  COLORS	##################
L_RED					=			\033[0;31m
L_REDB					=			\033[1;31m
L_WHITE					=			\033[0;37m
L_WHITEB				=			\033[1;37m
L_YELLOW				=			\033[0;33m
L_YELLOWB				=			\033[1;33m
L_GREEN					=			\033[0;32m
L_GREENB				=			\033[1;32m
################ CONFIG		##################
NAME					=			service
################ RULES		##################
all: $(NAME)

$(NAME):
	@go			build -o $(NAME)
	@echo		"$(L_GREENB)[Built Program]$(L_WHITE)"

fclean:
	@rm			-f $(NAME)
	@echo		"$(L_GREENB)[Program Executable Removed]$(L_WHITE)"

re: fclean all 