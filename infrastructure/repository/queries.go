package repository


var (
    CREATE_EVENEMENT_COMPLET_QUERY string = `
       SELECT creer_evenement_complet($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
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