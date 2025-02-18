#let header(members, guide) = {
align(center)[ 
  = Cobweb
  = A minimal website archiving solution
  \
  === Synopsis submitted to \ Shri Ramdeobaba College of Engineering & Management, Nagpur \ in partial fulfillment of requirement for the award of the degree of

  == Bachelor of Technology (B.Tech)
  \
  _*In*_
  \
  === COMPUTER SCIENCE AND ENGINEERING  (Cyber Security)
  === Semester - VI
  \
  *_By_*\ 
  \
  #for (name, roll) in members {
    [#name (#roll) \ ]
  }
  
  _*Guide*_ \ \
    #guide \ \

  #image("logo.png", fit: "contain", height: 15%)


  ==== Department of Computer Science and Engineering -- Cyber Security \ Shri Ramdeobaba College of Engineering & Management, Nagpur 440 013

  (An Autonomous Institute affiliated to Rashtrasant Tukdoji Maharaj \ Nagpur University Nagpur)

  *December 2024*
]
pagebreak()
}

#let footer(member, guide) = {
let (f1name, f1roll) = member.first()
  [
    #table(  
      columns: (1fr, 4fr, 4fr),
      table.header(
        [*Roll No.*], [*Name of Students*], [*Name of Guide*]
      ),
      align: center,

      [ #f1roll ],
      [ #f1name ],
      table.cell(
        [Firdous Sadaf],
        rowspan: 4,
        align: horizon,
      ),
      ..for (name, roll) in member.slice(1) {
            ([#roll], [#name])
      }
    )
    \ \ \
    *Approved by:*
    \ \ \
    #grid(
      rows: 3cm,
      columns: 2,
      gutter: 1fr,
      [*Head of Department* #parbreak() *Cyber Security*],
      [*Guide* #v(0pt) #guide ],
    )
  ]
}
