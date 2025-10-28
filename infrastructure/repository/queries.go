package repository


var (
    CREATE_EVENEMENT_COMPLET_QUERY string = `
        SELECT creer_evenement_complet(
            $1, $2, $3, $4, $5,  -- événement
            $6, $7, $8, $9,       -- lieu  
            $10::jsonb,           -- tarifs
            $11::jsonb            -- fichiers
        )
    `

    FIND_EVENEMENT_BY_ID_QUERY string = `
    SELECT 
        evenement_id,
        titre,
        description_evenement,
        date_debut,
        date_fin,
        type_evenement,
        lieu,
        tarifs,
        fichiers,
        statistiques
    FROM vue_evenement_complet 
    WHERE evenement_id = $1
`
)