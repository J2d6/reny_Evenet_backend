-- ================================================
-- VUE : Détails complets d'un événement
-- ================================================
CREATE OR REPLACE VIEW vue_evenement_details AS
SELECT 
    e.id as evenement_id,
    e.titre,
    e.description as description_evenement,
    e.date_debut,
    e.date_fin,
    
    -- Informations du type d'événement
    te.id as type_evenement_id,
    te.nom as type_evenement_nom,
    te.description as type_evenement_description,
    
    -- Informations du lieu
    l.id as lieu_id,
    l.nom as lieu_nom,
    l.adresse as lieu_adresse,
    l.ville as lieu_ville,
    l.capacite as lieu_capacite,
    
    -- Informations des tarifs
    t.id as tarif_id,
    t.prix,
    t.nombre_places as nombre_places_tarif,
    
    -- Informations du type de place
    tp.id as type_place_id,
    tp.nom as type_place_nom,
    tp.description as type_place_description,
    tp.avantages as type_place_avantages,
    
    -- Statistiques des places par tarif
    COUNT(p.id) as places_crees,
    COUNT(CASE WHEN p.etat_code = 'disponible' THEN 1 END) as places_disponibles,
    COUNT(CASE WHEN p.etat_code = 'reservee' THEN 1 END) as places_reservees,
    COUNT(CASE WHEN p.etat_code = 'vendue' THEN 1 END) as places_vendues,
    COUNT(CASE WHEN p.etat_code IN ('annulee', 'maintenance') THEN 1 END) as places_inactives

FROM evenement e
JOIN type_evenement te ON e.type_id = te.id
JOIN lieu l ON e.lieu_id = l.id
JOIN tarif t ON t.evenement_id = e.id
JOIN type_place tp ON t.type_place_id = tp.id
LEFT JOIN place p ON p.tarif_id = t.id

GROUP BY 
    e.id, e.titre, e.description, e.date_debut, e.date_fin,
    te.id, te.nom, te.description,
    l.id, l.nom, l.adresse, l.ville, l.capacite,
    t.id, t.prix, t.nombre_places,
    tp.id, tp.nom, tp.description, tp.avantages;