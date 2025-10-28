package reposiory


var (
    CREATE_EVENEMENT_COMPLET_QUERY string = `
        SELECT creer_evenement_complet(
            $1, $2, $3, $4, $5,  -- événement
            $6, $7, $8, $9,       -- lieu  
            $10::jsonb,           -- tarifs
            $11::jsonb            -- fichiers
        )
    `
)