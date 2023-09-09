package auth

import "github.com/nedpals/supabase-go"

func NewSupabaseAuthClient(supabaseUrl string, supabaseKey string) *supabase.Client {
	return supabase.CreateClient(supabaseUrl, supabaseKey)
}
