#!/usr/bin/env zsh

# Function to generate distinct random numbers
generate_random_numbers() {
    local count=$1
    local numbers=()
    while ((${#numbers[@]} < count)); do
        num=$((RANDOM % 100))
        if (( ! ${numbers[(Ie)$num]} )); then
            numbers+=($num)
        fi
    done
    echo "${numbers[*]}"
}

echo ""
echo "Try to run ./push-swap"
echo "Does it display nothing?"
echo ""

./push-swap

echo ""
echo "Try to run ./push-swap \"2 1 3 6 5 8\""
echo "Does it display a valid solution and less than 9 instructions?"
echo ""

./push-swap "2 1 3 6 5 8"

echo ""
./push-swap "2 1 3 6 5 8" | ../checker/checker "2 1 3 6 5 8"

echo ""
echo "Try to run ./push-swap \"0 1 2 3 4 5\""
echo "Does it display nothing?"
echo ""

./push-swap "0 1 2 3 4 5"

echo ""
echo "Try to run ./push-swap \"0 one 2 3\""
echo "Does it display the correct result as above? [Error]"
echo ""

./push-swap "0 one 2 3"

echo ""
echo "Try to run ./push-swap \"1 2 2 3\""
echo "Does it display the correct result as above? [Error]"
echo ""

./push-swap "1 2 2 3"

echo ""
echo "Try to run ./push-swap \"<5 random numbers>\" with 5 random numbers instead of the tag."
echo "Does it display a valid solution and less than 12 instructions?"
echo ""

random_numbers=$(generate_random_numbers 5)
echo "Random numbers: $random_numbers"
echo ""

./push-swap "$random_numbers"
echo ""
./push-swap "$random_numbers" | ../checker/checker "$random_numbers"

echo ""
echo "Try to run ./push-swap \"<5 random numbers>\" with 5 different random numbers instead of the tag."
echo "Does it still displays a valid solution and less than 12 instructions?"
echo ""

RANDOM=$$
random_numbers=$(generate_random_numbers 5)
echo "Random numbers: $random_numbers"
echo ""

./push-swap "$random_numbers"
echo ""
./push-swap "$random_numbers" | ../checker/checker "$random_numbers"

echo ""
echo "Try to run ./checker and input nothing."
echo "Does it display nothing?"
echo ""

../checker/checker

echo ""
echo "Try to run ./checker \"0 one 2 3\""
echo "Does it display the correct result as above? [Error]"
echo ""

../checker/checker "0 one 2 3"

echo ""
echo -E "Try to run echo -e \"sa\\npb\\nrrr\\n\" | ./checker \"0 9 1 8 2 7 3 6 4 5\""
echo "Does it display the correct result as above? [KO]"
echo ""

echo -e "sa\npb\nrrr\n" | ../checker/checker "0 9 1 8 2 7 3 6 4 5"

echo ""
echo -E "Try to run echo -e \"pb\nra\npb\nra\nsa\nra\npa\npa\n\" | ./checker \"0 9 1 8 2\""
echo "Does it display the correct result as above? [OK]"
echo ""

echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ../checker/checker "0 9 1 8 2"

echo ""
echo "Try to run ARG=\"4 67 3 87 23\"; ./push-swap \"\$ARG\" | ./checker \"\$ARG\""
echo "Does it display the correct result as above? [OK]"
echo ""

ARG="4 67 3 87 23"; ./push-swap "$ARG" | ../checker/checker "$ARG"

echo ""
echo "As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)"
echo ""

echo ""
echo "*** BONUS ***"
echo ""

echo "GENERAL:"

echo ""
echo "Try to run ARG=\"<100 random numbers>\"; ./push-swap \"\$ARG\" with 100 random different numbers instead of the tag."
echo "+Does it display less than 700 commands?"
echo ""

RANDOM=$$
random_numbers=$(generate_random_numbers 100)
echo "Random numbers: $random_numbers"

echo ""
ARG=$random_numbers; ./push-swap "$ARG"
echo ""
line_count=$(./push-swap "$random_numbers" | wc -l)
echo "Number of commands: $line_count"

echo ""
echo "Try to run ARG=\"<100 random numbers>\"; ./push-swap \"\$ARG\" | ./checker \"\$ARG\" with the same 100 random numbers as before instead of the tag."
echo "+Does it display the correct result as above? [OK]"
echo ""

ARG=$random_numbers; ./push-swap "$ARG" | ../checker/checker "$ARG"

echo ""
echo "BASIC:"

echo ""
echo "+Does the code obey the good practices?"
echo "+Is there a test file for this code?"
echo "+Are the tests checking each possible case?"

echo ""
echo "SOCIAL:"
echo ""

echo "+Did you learn anything from this project?"
echo "+Would you recommend/nominate this program as an example for the rest of the school?"
echo ""
